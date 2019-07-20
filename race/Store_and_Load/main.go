package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	//this is a flag to alert the running goroutines to shutdown
	shutdown int64

	//wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Create two goroutines
	go doWork("A")
	go doWork("B")

	// give the goroutines time to run
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	// wait for the goroutines to finish
	wg.Wait()

	//doWork simulatess a goroutuine performing work and checking the shutdown flag to terminate early
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		//Do we need to shutdown
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
