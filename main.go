package main

import (
	"log"
	"os"

	"./search"
)

func init() {
	//Change the device login to stdout
	// By default, the logger is set to write to the stderr device
	log.SetOutput(os.Stdout)
}

func main() {
	//Perform the search of the specified term:
	search.Run("president")
}
