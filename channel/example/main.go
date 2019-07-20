package main

import "fmt"

func main() {
	//Buffered channell of string
	//this buffer contains 10 values
	buffered := make(chan string, 10)

	// send a string through the channel
	buffered <- "Gopher"

	//receive a string from the channel
	value := <-buffered

	fmt.Println(value)
}
