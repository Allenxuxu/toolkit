package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	lock uintptr
}

func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapUintptr(&l.lock, 0, 1) {
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	atomic.StoreUintptr(&l.lock, 0)
}
