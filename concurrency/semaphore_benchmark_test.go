package concurrency

import (
	"sync/atomic"
	"testing"
)

func BenchmarkSemaphoreAcquireRelease(b *testing.B) {
	sem := NewSemaphore(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sem.Acquire()
		sem.Release()
	}
}

//func BenchmarkSemaphoreWait(b *testing.B) {
//	sem := NewSemaphore(10)
//
//	// Create `b.N` tasks
//	for i := 0; i < b.N; i++ {
//		sem.ProcessAndRelease(func() {
//			time.Sleep(1 * time.Millisecond) // Simulate work
//		})
//	}
//
//	b.ResetTimer()
//	sem.Wait()
//}

func BenchmarkSemaphoreProcessAndRelease(b *testing.B) {
	sem := NewSemaphore(10)
	var counter int32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sem.ProcessAndRelease(func() {
			atomic.AddInt32(&counter, 1)
			atomic.AddInt32(&counter, -1)
		})
	}
	sem.Wait()
}

func BenchmarkSemaphoreProcessAndReleaseReflect(b *testing.B) {
	sem := NewSemaphore(10)
	var counter int32

	increment := func(delta int32) {
		atomic.AddInt32(&counter, delta)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sem.ProcessAndReleaseReflect(increment, int32(1))
	}
	sem.Wait()
}

func BenchmarkSemaphoreHighConcurrency(b *testing.B) {
	sem := NewSemaphore(100) // High concurrency
	var counter int32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sem.ProcessAndRelease(func() {
			atomic.AddInt32(&counter, 1)
			atomic.AddInt32(&counter, -1)
		})
	}
	sem.Wait()
}
