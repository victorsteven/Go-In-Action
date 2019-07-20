package main

import (
	"fmt"

	"../counters"
)

func main() {
	// counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counters.New(10))

	theUser := counters.Admin{
		Rights: 10,
	}
	theUser.Name = "Mark" //WE could not do this in the inside the Admin struct, since the user is unexported
	theUser.Email = "mark@gmail.com"

	fmt.Println(theUser)

}
