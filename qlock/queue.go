package qlock

type queue struct {
	// dummy
}

func (q *queue) Add(threadID int) {}
func (q *queue) Remove() int      { return 0 }
func (q *queue) IsEmpty() bool    { return false }
