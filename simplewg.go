// Package simplewg provides a convenience wrapper for [sync.WaitGroup].
// It automates the invocation of sync.WaitGroup.Add and sync.WaitGroup.Done
package simplewg

import (
	"sync"
	"sync/atomic"
)

// Wg struct wraps sync.WaitGroup to offer enhanced concurrency control.
// It manages the lifecycle of goroutines and ensures synchronization between them.
type Wg struct {
	wg   sync.WaitGroup
	wait atomic.Pointer[chan struct{}]
	// Ensures the wait channel is initialized only once to reduce calls to make.
	waitOnce sync.Once
}

// Go method initiates a function in a separate goroutine.
// Once [Wg.Wait] has been called for the first time, subsequent calls to Go will not execute the function.
// Returns true if the function is executed, false otherwise.
func (wg *Wg) Go(fn func()) (ok bool) {
	if wg.wait.Load() == nil {
		wg.wg.Add(1)
		go func() {
			defer wg.wg.Done()
			fn()
		}()
		return true
	}
	return false
}

// Wait method blocks the caller until all goroutines have completed.
// It is designed to handle multiple concurrent calls.
func (wg *Wg) Wait() { <-wg.Done() }

// Done returns a channel that gets closed when all goroutines have finished.
// The first call to Done prevents new goroutines from starting via Go.
func (wg *Wg) Done() <-chan struct{} {
	var c chan struct{}
	wg.waitOnce.Do(func() { c = make(chan struct{}) })
	if wg.wait.CompareAndSwap(nil, &c) {
		go func() {
			defer close(c)
			wg.wg.Wait()
		}()
	}
	return *wg.wait.Load()
}
