package ch29

import "sync"

const (
	NumCPUs = 6
)

type ApproxCounter struct {
	global    int
	gmu       sync.Mutex
	local     []int
	lmu       []sync.Mutex
	threshold int
}

func NewApproxCounter(threshold int) *ApproxCounter {
	return &ApproxCounter{
		threshold: threshold,
		local:     make([]int, NumCPUs),
		lmu:       make([]sync.Mutex, NumCPUs),
	}
}

func (ac *ApproxCounter) Get() int {
	ac.gmu.Lock()
	defer ac.gmu.Unlock()
	return ac.global
}

func (ac *ApproxCounter) Update(threadID, amount int) {
	if threadID > len(ac.lmu) || threadID < 0 {
		panic("threadID is out of bounds")
	}
	ac.lmu[threadID].Lock()
	defer ac.lmu[threadID].Unlock()
	ac.local[threadID] += amount
	if ac.local[threadID] >= ac.threshold {
		ac.gmu.Lock()
		defer ac.gmu.Unlock()
		ac.global += ac.local[threadID]
		ac.local[threadID] = 0
	}
}

func (ac *ApproxCounter) UpdateNoLocalLock(threadID, amount int) {
	ac.local[threadID] += amount
	if ac.local[threadID] >= ac.threshold {
		ac.gmu.Lock()
		defer ac.gmu.Unlock()
		ac.global += ac.local[threadID]
		ac.local[threadID] = 0
	}
}
