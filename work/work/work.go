//package work manages a pool of goroutines to perform work
package work

import "sync"

//  Worker must be implemnted by types taht want to use the work pool

type Worker interface {
	Task()
}

//Pool provides a pool of goroutines that can execute any worker tasks that are submitted

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run submit work to the poop
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
