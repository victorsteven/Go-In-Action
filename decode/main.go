package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "Steven", "Text": "Hello bro"}
	{"Name": "Emma", "Text": "Hello man"}
	{"Name": "Mark", "Text": "Hello boss"}
	`

	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
