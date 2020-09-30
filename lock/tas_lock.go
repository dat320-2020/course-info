package lock

// TestAndSet takes a pointer to an int and a new value
// to be set at the location pointed to by oldPtr.
// TestAndSet returns the value of the int pointed to by
// oldPtr (before updating its value).
// This is pseudo code for a CPU instruction performed atomically.
func TestAndSet(oldPtr *int, new int) int {
	old := *oldPtr
	*oldPtr = new
	return old
}

func (l *lock) TASLock() {
	for TestAndSet(&l.flag, 1) == 1 {
		// spin-wait
	}
}
