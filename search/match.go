package search

import (
	"fmt"
	"log"
)

//Result contains the result the search
type Result struct {
	Field   string
	Content string
}

//Matcher defines the behaviour required by types that want to implement a new search type.

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//MAtch is launched as a goroutine for each individual feed to run searches concurrently

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//Performthe search againt the specified matcher.

	searchResults, err := matcher.Search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	//Write the results to the channel
	for _, result := range searchResults {
		results <- result
	}
}

//Display writes results to the terminal window as thery are received by the individual goroutines

func Display(results chan *Result) {
	//The channel blocks until a result is written to the channel.
	//once the channel is closed the for loop terminates
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
