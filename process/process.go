package process

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

const ParallelLimit = 1024

// Parallel provides an abstraction on a key-value persistent database, which
// prevents functions from being executed twice, ever.
type Parallel struct {
	sm *semaphore.Weighted
	wg sync.WaitGroup
	er atomic.Value
}

func NewParallel() *Parallel {
	var sema = semaphore.NewWeighted(ParallelLimit)
	return &Parallel{sm: sema}
}

// DoAsync runs the function asynchronously. The given callback will be passed
// the given key string. This prevents common bugs with range. If the callback
// returns false, then it is not recorded.
func (p *Parallel) DoAsync(key, value string, fn func(k, v string) error) {
	// Skip if there's an error.
	if p.er.Load() != nil {
		return
	}

	p.wg.Add(1)
	p.sm.Acquire(context.Background(), 1)
	go func() {
		if err := fn(key, value); err != nil {
			p.er.Store(err)
		}
		p.sm.Release(1)
		p.wg.Done()
	}()
}

func (p *Parallel) Wait() {
	p.wg.Wait()
}

func (p *Parallel) Err() error {
	if err, ok := p.er.Load().(error); ok {
		return err
	}
	return nil
}
