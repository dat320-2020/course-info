package ch29_test

import (
	"course-info/ch29"
	"sync"
	"testing"
)

func TestConcurrentLinkedList(t *testing.T) {
	list := &ch29.ListL{}
	elements := 10_000
	var wg sync.WaitGroup
	wg.Add(ch29.NumCPUs)
	for i := 0; i < ch29.NumCPUs; i++ {
		go func() {
			for k := 0; k < elements; k++ {
				e := k + elements*ch29.NumCPUs
				list.InsertL(e)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	wg.Add(ch29.NumCPUs)
	for i := 0; i < ch29.NumCPUs; i++ {
		go func() {
			for k := 0; k < elements; k++ {
				e := k + elements*ch29.NumCPUs
				elm := list.LookupL(e)
				if elm == nil {
					t.Errorf("Lookup(%d) = %v, expected %d", e, elm, e)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
