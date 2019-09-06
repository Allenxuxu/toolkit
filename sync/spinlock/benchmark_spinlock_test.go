package spinlock

import (
	"sync"
	"testing"
)

func BenchmarkSpinLock(b *testing.B) {
	var mu SpinLock

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			mu.Unlock() // nolint
		}
	})
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			mu.Unlock() // nolint
		}
	})
}
