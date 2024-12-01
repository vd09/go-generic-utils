package concurrency

import (
	"sync"
)

// Task is a type that represents a function to be executed by a worker.
type Task func()

// WorkerPool is a struct that manages a pool of workers to execute tasks concurrently.
type WorkerPool struct {
	tasks      chan Task
	wg         sync.WaitGroup
	numWorkers int
}

// NewWorkerPool creates a new WorkerPool with a specified number of workers and a maximum number of tasks in the queue.
func NewWorkerPool(numWorkers int, maxTasks int) *WorkerPool {
	pool := &WorkerPool{
		tasks:      make(chan Task, maxTasks),
		numWorkers: numWorkers,
		wg:         sync.WaitGroup{},
	}

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

// worker is a function that is executed by each worker goroutine. It processes tasks from the tasks channel.
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for task := range wp.tasks {
		if task != nil {
			task()
		}
	}
}

// AddTask adds a new task to the worker pool for execution.
func (wp *WorkerPool) AddTask(task Task) {
	wp.tasks <- task
}

// Wait blocks until all tasks have been completed and all workers have stopped.
func (wp *WorkerPool) Wait() {
	close(wp.tasks) // Signal to the workers that no more tasks will be sent
	wp.wg.Wait()    // Wait for all workers to finish
}
