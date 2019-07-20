package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup

	//

)

func main() {
	wg.Add(2)

	//Create two goroutines
	go intCounter(1)
	go intCounter(2)

	//wait for the goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}

func intCounter(id int) {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		//Yield the thread and be placed back in queue
		runtime.Gosched()

		//Increment our local value of counter
		value++

		//Store the value back into counter
		counter = value
	}
}
