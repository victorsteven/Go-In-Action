package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact represents our JSON string

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

//JSON contains a sample string to unmarshal

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "3435.435345.54",
		"cell": "253514.45345.325"
	} 
}`

func main() {
	// Unmarshal the JSON string into our variable

	// decoding or unmarshalling into c of type struct:
	// var c Contact
	// err := json.Unmarshal([]byte(JSON), &c)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	return
	// }

	// decoding or unmarshalling into a map:
	// with a key of string a value of interface
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	//Because the value for each key is of type interface{}, you need to convert the value to the proper native type in order to work with the value.
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])

	fmt.Println(c)
}
