package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//Allocate 2 logical processor for the scheduler to use
	runtime.GOMAXPROCS(2)
	//wg is used to wait for the program to finish
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//Declare an anonymous function and create a goroutine
	go func() {
		//Schedule the call to Done to tell main when we are done
		defer wg.Done()

		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	//Declare an anonymous function and create a goroutine
	go func() {
		//Schedule the call to Done to tell the main we are done
		defer wg.Done()

		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	//Wait the goroutine to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}