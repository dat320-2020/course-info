package ch29

import "sync"

type counter struct {
	val int
	mu  sync.Mutex
}

func (c *counter) NoLockInc() {
	c.val++
}

func (c *counter) NoLockDec() {
	c.val--
}

func (c *counter) NoLockGet() int {
	return c.val
}
