package locks

import (
	"sync"
	"testing"
)

func TestReentrantLock(t *testing.T) {
	cnt := 0
	reentrantLock := &ReentrantLock{}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		reentrantLock.Lock()
		cnt++
		reentrantLock.Lock()
		cnt++
		reentrantLock.Unlock()
		reentrantLock.Unlock()
		wg.Done()
	}()
	reentrantLock.Lock()
	cnt--
	reentrantLock.Unlock()
	wg.Wait()
	if cnt != 1 {
		t.Errorf("expected: %v, got: %v", 1, cnt)
	}
}
