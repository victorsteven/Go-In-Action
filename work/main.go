package main

import (
	"log"
	"sync"
	"time"

	"./work"
)

var mames = []string{
	"steven",
	"bob",
	"ade",
}

// namePrinter provides special support for printing names
type namePrinter struct {
	name string
}

// Task impelements the Worker interface
func (m *namePrinter) Tasks() {
	log.Println(m.name)
	time.Sleep(time.Second)

}

func main() {
	// Create a work pool with 2 goroutines
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// iterate over the slice of names
		for _, name := range names {
			// Create a namePrinter and provide the specific name
			np := namePrinter{
				name: name,
			}
		}
	}
}
