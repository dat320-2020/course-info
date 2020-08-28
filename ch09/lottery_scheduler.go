package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type lotteryScheduler interface {
	Schedule()
	Add(job)
}

type jobList struct {
	head *job
}

type job struct {
	name    string
	tickets int
	next    *job
}

func (l jobList) String() string {
	if l.head == nil {
		return "No jobs in job list\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Job: %s, Tickets: %d\n", l.head.name, l.head.tickets)
	tmp := l.head
	for tmp.next != nil {
		fmt.Fprintf(&b, "Job: %s, Tickets: %d\n", tmp.next.name, tmp.next.tickets)
		tmp = tmp.next
	}
	return fmt.Sprintf("Jobs: \n%v", b.String())
}

func (l *jobList) Add(j *job) {
	if l.head == nil {
		l.head = j
		return
	}
	tmp := l.head
	for tmp != nil {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = j
}

func (l jobList) Schedule() {
	var counter int32
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Int31n(399)
	fmt.Println("rnd=", rnd)
	current := l.head
	for current != nil {
		counter += int32(current.tickets)
		if counter > rnd {
			break
		}
		current = current.next
	}
	fmt.Println("Winner is: ", current.name)
}

func main() {
	a := &job{name: "A", tickets: 100}
	b := &job{name: "B", tickets: 50}
	c := &job{name: "C", tickets: 250}
	// var lottery lotteryScheduler
	lottery := &jobList{}
	lottery.Add(a)
	lottery.Add(b)
	lottery.Add(c)
	fmt.Println(lottery)
	lottery.Schedule()
}
