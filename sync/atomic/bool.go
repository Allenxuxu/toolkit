package atomic

import "sync/atomic"

type Bool struct {
	b int32
}

func New(b bool) *Bool {
	var ret Bool
	if b {
		ret.b = 1
	}

	return &ret
}

func (a *Bool) Set(b bool) bool {
	var newV int32
	if b {
		newV = 1
	}
	return atomic.SwapInt32(&a.b, newV) == 1
}

func (a *Bool) Get() bool {
	return atomic.LoadInt32(&a.b) == 1
}

func (a *Bool) CompareAndSwap(old, new bool) bool {
	var o, n int32

	if old {
		o = 1
	}
	if new {
		n = 1
	}

	return atomic.CompareAndSwapInt32(&a.b, o, n)
}
