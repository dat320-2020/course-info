package safeint_test

import (
	"course-info/safeint"
	"testing"
)

func TestSetGet(t *testing.T) {
	i := &safeint.SharedInt{}
	got := i.Get()
	want := 0
	if got != want {
		t.Errorf("i.Val = %d, want %d", got, want)
	}
	i.Set(42)
	got = i.Get()
	want = 42
	if got != want {
		t.Errorf("i.Get() = %d, want %d", got, want)
	}
}

func TestConcurrentIncrement(t *testing.T) {
	t.Parallel()
	i := &safeint.SharedInt{}
	expected := 1000000
	ch := make(chan bool)
	g := func() {
		for k := 0; k < expected; k++ {
			i.Inc()
		}
		ch <- true
	}
	go g()
	go g()
	go g()
	<-ch
	<-ch
	<-ch
	got := i.Get()
	want := 3 * expected
	if got != want {
		t.Errorf("i.Get() = %d, want %d", got, want)
	}
}
