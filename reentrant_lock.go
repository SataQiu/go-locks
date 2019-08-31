package locks

import (
	"sync"
)

type ReentrantLock struct {
	mux sync.Mutex
	c   chan struct{}
	v   int32  // store lock times
	gid uint64 // store goroutine id
}

func (rl *ReentrantLock) Lock() {
	gid := getGID()
	for {
		rl.mux.Lock()
		if rl.c == nil {
			rl.c = make(chan struct{}, 1)
		}
		if rl.v == 0 || rl.gid == gid {
			rl.v++
			rl.gid = gid
			rl.mux.Unlock()
			break
		}
		rl.mux.Unlock()
		<-rl.c
	}
}

func (rl *ReentrantLock) Unlock() {
	rl.mux.Lock()
	if rl.c == nil {
		rl.c = make(chan struct{}, 1)
	}
	if rl.v <= 0 {
		rl.mux.Unlock()
		panic("Unable to Unlock, not locked.")
	}
	rl.v--
	if rl.v == 0 {
		rl.gid = 0
	}
	rl.mux.Unlock()
	select {
	case rl.c <- struct{}{}:
	default:
	}
}
