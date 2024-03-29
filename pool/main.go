package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"./pool"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

// dbConnection simulates a resource to share
type dbConnection struct {
	ID int32
}

// Close implements the io.Closer interface so dbConnection can be managed by the pool. close performs any resource release management

func (dbConn *dbConnection) Close() error {
	log.Println("Close:  Connection", dbConn.ID)
	return nil
}

// idCounter provides support for giving each connection a unique id
var idCounter int32

// createConnection is a factory methos that will be called by the pool when a new connection is needed

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	//Create the pool to manage our connection
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	//Perform queries using connections from the pool

	for query := 0; query < maxGoroutines; query++ {
		// Each goroutines needs its own copy of the query value else they will all be sharing the same query variable

		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	//Wait for the goroutines to finish
	wg.Wait()

	// Close the pool
	log.Println("Shutdown program.")
	p.Close()
}

//performQueries tests the resources pool of connections
func performQueries(query int, p *pool.Pool) {
	// Acquire a connection from the pool
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	//Release the connection back to the pool:
	defer p.Release(conn)

	//Wait to simulate a query response:
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
