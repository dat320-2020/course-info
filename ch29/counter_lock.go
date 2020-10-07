package ch29

func (c *counter) Inc() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func (c *counter) IncWithDefer() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *counter) Dec() {
	c.mu.Lock()
	c.val--
	c.mu.Unlock()
}

func (c *counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
