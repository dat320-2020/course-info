package qlock

import "course-info/lock"

type queueLock struct {
	flag  int
	guard int
	queue *queue
}

func (l *queueLock) QLock() {
	for lock.TestAndSet(&l.guard, 1) == 1 {
		// spin
	}
	if l.flag == 0 {
		l.flag = 1
		l.guard = 0
	} else {
		l.queue.Add(getThreadID())
		l.guard = 0
		park() // link yield
	}
}

func (l *queueLock) QUnlock() {
	for lock.TestAndSet(&l.guard, 1) == 1 {
		// spin
	}
	if l.queue.IsEmpty() {
		l.flag = 0
	} else {
		threadID := l.queue.Remove()
		unpark(threadID)
	}
	l.guard = 0
}

func getThreadID() int    { return 0 }
func park()               {}
func unpark(threadID int) {}
