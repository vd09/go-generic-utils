package concurrency

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWorkerPool(t *testing.T) {
	numWorkers := 5
	maxTasks := 10

	wp := NewWorkerPool(numWorkers, maxTasks)
	assert.NotNil(t, wp, "NewWorkerPool should return a non-nil WorkerPool")
	assert.Equal(t, numWorkers, wp.numWorkers, "WorkerPool should have the correct number of workers")

	// Add a dummy task and ensure the worker processes it
	done := make(chan struct{})
	wp.AddTask(func() {
		close(done)
	})

	wp.Wait()
	assert.NotPanics(t, func() { <-done }, "WorkerPool should process tasks correctly")
}

func TestWorkerPoolAddTask(t *testing.T) {
	numWorkers := 3
	maxTasks := 5
	wp := NewWorkerPool(numWorkers, maxTasks)

	var counter int32

	// Add tasks
	for i := 0; i < 10; i++ {
		wp.AddTask(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(50 * time.Millisecond) // Simulate work
		})
	}

	wp.Wait()

	assert.Equal(t, int32(10), counter, "All tasks should be completed")
}

func TestWorkerPoolConcurrency(t *testing.T) {
	numWorkers := 4
	maxTasks := 8
	wp := NewWorkerPool(numWorkers, maxTasks)

	var counter int32
	var maxConcurrent int32

	concurrencyTracker := int32(0)

	// Add tasks that track concurrency
	for i := 0; i < 20; i++ {
		wp.AddTask(func() {
			current := atomic.AddInt32(&concurrencyTracker, 1)
			atomic.AddInt32(&counter, 1)
			time.Sleep(50 * time.Millisecond)
			atomic.AddInt32(&concurrencyTracker, -1)

			// Track the maximum concurrency
			if current > atomic.LoadInt32(&maxConcurrent) {
				atomic.StoreInt32(&maxConcurrent, current)
			}
		})
	}

	wp.Wait()

	assert.Equal(t, int32(20), counter, "All tasks should be completed")
	assert.Equal(t, int32(numWorkers), maxConcurrent, "Maximum concurrency should equal the number of workers")
}

func TestWorkerPoolGracefulShutdown(t *testing.T) {
	numWorkers := 3
	maxTasks := 5
	wp := NewWorkerPool(numWorkers, maxTasks)

	// Add a few tasks
	var counter int32
	for i := 0; i < 5; i++ {
		wp.AddTask(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(50 * time.Millisecond)
		})
	}

	// Wait for tasks to complete
	wp.Wait()

	// Verify no tasks are pending, and workers have stopped
	assert.Equal(t, int32(5), counter, "All tasks should be completed before shutdown")
}

func TestWorkerPoolZeroWorkers(t *testing.T) {
	numWorkers := 0
	maxTasks := 5
	wp := NewWorkerPool(numWorkers, maxTasks)

	var executed bool

	// Add a task
	wp.AddTask(func() {
		executed = true
	})

	// Wait and verify no tasks were processed
	wp.Wait()
	assert.False(t, executed, "No tasks should be processed when there are zero workers")
}

func TestWorkerPoolLargeTaskQueue(t *testing.T) {
	numWorkers := 3
	maxTasks := 100
	wp := NewWorkerPool(numWorkers, maxTasks)

	var counter int32

	// Add a large number of tasks
	for i := 0; i < 100; i++ {
		wp.AddTask(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	wp.Wait()

	assert.Equal(t, int32(100), counter, "All tasks should be completed even with a large task queue")
}
