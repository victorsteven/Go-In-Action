package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create an unbeffered channel
	court := make(chan int)

	// Add a count of two, one for each goroutine
	wg.Add(2)

	//Launch two players
	go player("Nadal", court)
	go player("Victor", court)

	//Start the set
	court <- 1

	// Wait for the game to finish
	wg.Wait()

}

// Player simulates a person playing the game of tennis
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us:
		ball, ok := <-court
		if !ok {
			// if the channel was closed, we won
			fmt.Printf("Player %s Won\n", name)
			return
		}
		//pick a random number and see if we miss the ball
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal that we lost
			close(court)
			return
		}

		// Display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player
		court <- ball
	}
}
