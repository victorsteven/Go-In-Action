package main

import "fmt"

func main() {

	//Key of string and value of int
	// dict := make(map[string]int)
	//or using string literal
	// dict := map[string]int{}

	//create a map using a slice of strings as the key
	// dict := map[[]string]int{}
	// Compiler Exception:
	// invalid map key type []string
	// There’s nothing stopping you from using a slice as a map value. This can come in handy when you need a single map key to be associated with a collection of data.

	//uzing a slice of strings as value:
	// dict := map[int][]string{}

	//Creating  a nil map
	// var colors map[string]string
	// colors["Red"] = "red"
	// fmt.Println(colors) //this prints an error

	// Create an map to store colors and their color codes:
	colors2 := map[string]string{}
	colors2["Blue"] = "blue"
	fmt.Println(colors2) //this prints well

	//Retrieving a value from a map and testing the existence
	value, exists := colors2["Blue"]
	//did the key exists ?
	if exists {
		fmt.Println(value)
	}
	//second way:
	if value != "" {
		fmt.Println(value)
	}

	//adding the the map:
	colors2["White"] = "white"

	// deleting from the map:
	delete(colors2, "Blue")
	fmt.Println(colors2)

	//calling the delete function:
	deleteItem(colors2, "White")

	fmt.Println(colors2)

}

//we can write a function that delete a key from the map:
func deleteItem(colors2 map[string]string, key string) {
	delete(colors2, key)
}
