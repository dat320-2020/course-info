package ch29_test

import (
	"course-info/ch29"
	"fmt"
	"sync"
	"testing"
)

func BenchmarkConcurrentApproxCounter(b *testing.B) {
	for threshold := 1; threshold < 2048; threshold = 2 * threshold {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				runLoop(threshold)
			}
		})
	}
}

func BenchmarkConcurrentApproxCounterNoLocalLock(b *testing.B) {
	for threshold := 1; threshold < 2048; threshold = 2 * threshold {
		b.Run(fmt.Sprintf("%5d", threshold), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				runLoopNoLocalLock(threshold)
			}
		})
	}
}

const totalCount = 1_000_000

func runLoop(threshold int) {
	c := ch29.NewApproxCounter(threshold)
	var wg sync.WaitGroup
	wg.Add(ch29.NumCPUs)
	for i := 0; i < ch29.NumCPUs; i++ {
		go func(threadID int) {
			for k := 0; k < totalCount; k++ {
				c.Update(threadID, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func runLoopNoLocalLock(threshold int) {
	c := ch29.NewApproxCounter(threshold)
	var wg sync.WaitGroup
	wg.Add(ch29.NumCPUs)
	for i := 0; i < ch29.NumCPUs; i++ {
		go func(threadID int) {
			for k := 0; k < totalCount; k++ {
				c.UpdateNoLocalLock(threadID, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
