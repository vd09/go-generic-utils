package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSemaphore(t *testing.T) {
	maxGoroutines := 5
	s := NewSemaphore(maxGoroutines)

	assert.NotNil(t, s, "NewSemaphore should return a non-nil Semaphore")
	assert.Equal(t, maxGoroutines, s.maxGoroutines, "Semaphore should have the correct maxGoroutines")
	assert.Equal(t, 0, len(s.sem), "Semaphore channel should be empty initially")
}

func TestSemaphoreAcquireRelease(t *testing.T) {
	maxGoroutines := 3
	s := NewSemaphore(maxGoroutines)

	// Acquire slots
	for i := 0; i < maxGoroutines; i++ {
		s.Acquire()
		assert.Equal(t, i+1, len(s.sem), "Semaphore channel should reflect acquired slots")
	}

	// Release slots
	for i := 0; i < maxGoroutines; i++ {
		s.Release()
		assert.Equal(t, maxGoroutines-i-1, len(s.sem), "Semaphore channel should reflect released slots")
	}
}

func TestSemaphoreWait(t *testing.T) {
	s := NewSemaphore(2)
	var counter int32

	for i := 0; i < 5; i++ {
		s.ProcessAndRelease(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(100 * time.Millisecond)
		})
	}

	s.Wait()
	assert.Equal(t, int32(5), counter, "All goroutines should complete before Wait returns")
}

func TestSemaphoreProcessAndRelease(t *testing.T) {
	s := NewSemaphore(2)
	var counter int32

	// Process two tasks concurrently
	for i := 0; i < 4; i++ {
		s.ProcessAndRelease(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(100 * time.Millisecond)
		})
	}

	s.Wait()
	assert.Equal(t, int32(4), counter, "All tasks should complete and increment the counter")
}

func TestSemaphoreProcessAndReleaseReflect(t *testing.T) {
	s := NewSemaphore(3)
	var results []string
	var mu sync.Mutex

	appendToResults := func(prefix string, value int) {
		mu.Lock()
		defer mu.Unlock()
		results = append(results, fmt.Sprintf("%s%d", prefix, value))
	}

	for i := 0; i < 5; i++ {
		s.ProcessAndReleaseReflect(appendToResults, "test", i)
	}

	s.Wait()
	assert.Equal(t, 5, len(results), "All tasks should complete and populate the results slice")
	for i := 0; i < 5; i++ {
		assert.Contains(t, results, fmt.Sprintf("test%d", i), "Results should contain the expected values")
	}
}

func TestSemaphoreConcurrentExecutionLimit(t *testing.T) {
	maxGoroutines := 2
	s := NewSemaphore(maxGoroutines)
	var counter int32

	// Start multiple tasks
	for i := 0; i < 10; i++ {
		s.ProcessAndRelease(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(100 * time.Millisecond)
			atomic.AddInt32(&counter, -1)
		})

		assert.LessOrEqual(t, int(atomic.LoadInt32(&counter)), maxGoroutines, "Concurrent execution should not exceed maxGoroutines")
	}

	s.Wait()
	assert.Equal(t, int32(0), counter, "Counter should return to 0 after all tasks complete")
}
