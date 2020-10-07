package ch29

import (
	"sync"
	"testing"
)

const (
	totalCount = 1_000_000
	numCPUs    = 6
)

func TestCounter(t *testing.T) {
	count := &counter{}
	count.Inc()
	got := count.Get()
	if got != 1 {
		t.Errorf("count.Get() = %d, expected 1", got)
	}
}

func TestConcurrentCounter(t *testing.T) {
	count := &counter{}
	var wg sync.WaitGroup
	for i := 0; i < totalCount; i++ {
		wg.Add(1)
		go func() {
			count.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	got := count.Get()
	if got != totalCount {
		t.Errorf("count.Get() = %d, expected %d", got, totalCount)
	}
}

func BenchmarkNoLockCounter(b *testing.B) {
	// does not use goroutines or locks
	count := &counter{}
	tc := totalCount * numCPUs
	for k := 0; k < b.N; k++ {
		for j := 0; j < tc; j++ {
			count.NoLockInc()
		}
	}
}

func BenchmarkConcurrentCounter(b *testing.B) {
	count := &counter{}
	var wg sync.WaitGroup
	for k := 0; k < b.N; k++ {
		wg.Add(numCPUs)
		for j := 0; j < numCPUs; j++ {
			go func() {
				for i := 0; i < totalCount; i++ {
					count.Inc()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentNoLockCounter(b *testing.B) {
	count := &counter{}
	var wg sync.WaitGroup
	for k := 0; k < b.N; k++ {
		wg.Add(numCPUs)
		for j := 0; j < numCPUs; j++ {
			go func() {
				for i := 0; i < totalCount; i++ {
					count.NoLockInc()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
