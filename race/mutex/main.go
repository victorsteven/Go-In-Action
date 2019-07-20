package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines
	counter int

	wg sync.WaitGroup

	// mutex is used to define a critical section of  the code
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)

}

// incCounter increments the package level Counter variable using the Mutex to synchronize and provide safe access

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// onlu allow one goroutinee through this critical section at a time
		mutex.Lock()
		{
			//capture the value of counter

			value := counter
			//yield the thread and be placed back in queue
			runtime.Gosched()

			//increment our local value of counter
			value++

			// store the value back into counter
			counter = value
		}
		mutex.Unlock()
		//Release the lock and allow goroutine through
	}
}
