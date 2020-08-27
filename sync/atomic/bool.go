package atomic

import "sync/atomic"

type Bool struct {
	b int32
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
