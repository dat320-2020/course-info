package safeint

import "sync"

type SharedInt struct {
	mu  sync.Mutex
	val int
}

func (s *SharedInt) Set(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val = v
}

func (s *SharedInt) Get() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.val
}

func (s *SharedInt) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val++
}
