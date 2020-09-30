package lock

type lock struct {
	flag int
}

// Lock implements locking via load/store to memory.
// This implementation would violate correctness;
// see example trace in lock_trace.txt.
func (l *lock) Lock() {
	for l.flag == 1 {
		// spin-wait (do nothing)
	}
	l.flag = 1
}

func (l *lock) Unlock() {
	l.flag = 0
}
