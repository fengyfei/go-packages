package runner

import (
	"errors"
	"sync"
)

var (
	errRunnerClosed = errors.New("runner: runner closed")
)

type chanRoutineInterface interface {
	Run(<-chan bool)
}

type chanCloseGoRoutine func(<-chan bool)

func (f chanCloseGoRoutine) Run(close <-chan bool) {
	f(close)
}

// ChanCloserRunner stops a group of goroutines by closing a channel.
type ChanCloserRunner struct {
	mu        sync.Mutex
	waitGroup sync.WaitGroup
	close     chan bool
	closed    bool
}

// NewChanCloserRunner creates a runner instance.
func NewChanCloserRunner() *ChanCloserRunner {
	return &ChanCloserRunner{
		close: make(chan bool),
	}
}

// Go starts a goroutine.
func (r *ChanCloserRunner) Go(goroutine func(<-chan bool)) error {
	r.mu.Lock()

	if r.closed {
		r.mu.Unlock()
		return errRunnerClosed
	}

	r.waitGroup.Add(1)

	Go(func() {
		goroutine(r.close)
		r.waitGroup.Done()
	})

	r.mu.Unlock()

	return nil
}

// Interrupt stops all running goroutines.
func (r *ChanCloserRunner) Interrupt() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.closed {
		close(r.close)
	}

	r.waitGroup.Wait()

	r.closed = true
}

// Wait until all running goroutines to finish.
func (r *ChanCloserRunner) Wait() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.waitGroup.Wait()
}
