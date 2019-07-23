package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//Pool manages a set of resources that can be shared safely by multiple goroutines. the resources being shared must implement the io.Closer interface

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed is returned when an Acquire returns on a
// closed pool.

var ErrPoolClosed = errors.New("Pool has been closed")

//New creates a pool that manages resources. A pool requires a function that can allocate a new resource and the size of the pool

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

//Acquire retrieves a resource from the pool
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//check for a free resource
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared resources")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		//provide a resoiurce since there are none available
	default:
		log.Println("Acquirea: ", "New Resource")
		return p.factory()
	}
}

//Release places a new resources onto the pool
func (p *Pool) Release(r io.Closer) {
	//secure this operation with the Close operation
	p.m.Lock()
	defer p.m.Unlock()

	//if the pool is closed, discard the resource
	if p.closed {
		r.Close()
		return
	}

	select {
	//Attempt to place the new resourace on the queue
	case p.resources <- r:
		log.Println("Release:", "In Queue")
		//if the queue is already at capacity we close the resources
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources:
func (p *Pool) Close() {
	//secure this operation with the Release operation
	p.m.Lock()
	defer p.m.Unlock()

	//if the pool is alredy closed, dont do anything
	if p.closed {
		return
	}

	//set the pool as closed
	p.closed = true

	//close the channel before we drain the channel of its resources. if we dont do this, we will have a deadlock

	close(p.resources)

	//close the resources

	for r := range p.resources {
		r.Close()
	}
}
