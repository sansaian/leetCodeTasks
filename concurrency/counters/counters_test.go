package counters

import (
	"sync"
	"testing"
)

func Benchmark_Mutex(b *testing.B) {

	counterMutex := &counterWithMutex{
		mu:    sync.Mutex{},
		value: 0,
	}

	for n := 0; n < b.N; n++ {
		changeCounters(counterMutex, 10)
	}
}

func Benchmark_RWMutex(b *testing.B) {

	counterMutex := &counterWithRWMutex{
		rwMu:  sync.RWMutex{},
		value: 0,
	}

	for n := 0; n < b.N; n++ {
		changeCounters(counterMutex, 10)
	}
}

func Benchmark_Atomic(b *testing.B) {

	counterMutex := &counterWithAtomic{
		value: 0,
	}

	for n := 0; n < b.N; n++ {
		changeCounters(counterMutex, 10)
	}
}
