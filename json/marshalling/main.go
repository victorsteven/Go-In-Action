package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// Create a map of key/value pairs

	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "43535.4535.435",
		"cell": "3545.3434525.345",
	}

	//  Marshal the map into a JSON string
	// it returns a byte slice and an error
	//func MarshalIndent(v interface{}, prefix, indent string)
	// ([]byte, error) {}

	// data, err := json.MarshalIndent(c, "", "    ")
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
	// fmt.Println("this is the print")

}
