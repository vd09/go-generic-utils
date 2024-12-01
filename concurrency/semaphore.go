package concurrency

import (
	"reflect"
	"sync"
)

// Semaphore is a struct that encapsulates a semaphore pattern for concurrency control.
type Semaphore struct {
	maxGoroutines int
	sem           chan struct{}
	wg            sync.WaitGroup
}

// NewSemaphore creates a new Semaphore with a specified maximum number of concurrent goroutines.
func NewSemaphore(maxGoroutines int) *Semaphore {
	return &Semaphore{
		maxGoroutines: maxGoroutines,
		sem:           make(chan struct{}, maxGoroutines),
	}
}

// Acquire acquires a semaphore slot. Blocks if no slots are available.
func (s *Semaphore) Acquire() {
	s.wg.Add(1)
	s.sem <- struct{}{}
}

// Release releases a semaphore slot.
func (s *Semaphore) Release() {
	<-s.sem
	s.wg.Done()
}

// Wait waits for all goroutines to complete.
func (s *Semaphore) Wait() {
	s.wg.Wait()
}

// ProcessAndReleaseReflect processes a given function with arguments and releases the semaphore.
func (s *Semaphore) ProcessAndReleaseReflect(fn interface{}, args ...interface{}) {
	s.Acquire()
	go func() {
		defer s.Release()

		// Use reflection to call the function with the provided arguments
		fnValue := reflect.ValueOf(fn)
		fnArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			fnArgs[i] = reflect.ValueOf(arg)
		}

		fnValue.Call(fnArgs)
	}()
}

// ProcessAndRelease processes a given function with arguments and releases the semaphore.
func (s *Semaphore) ProcessAndRelease(fn func()) {
	s.Acquire()
	go func() {
		defer s.Release()
		fn()
	}()
}
