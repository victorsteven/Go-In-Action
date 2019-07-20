package main

import "fmt"

func main() {
	//Create a slice that have a length and capacity of 5 elements
	// slice := make([]string, 5)

	//A slice with a length of 3 and capacity of 5
	// slice := make([]int, 3, 5)

	//Declaring a slice with index positions
	//Initialize the 100th element with an empty string
	// slice := []string{99: ""}
	// fmt.Println(slice)

	// Create a slice with length and capacity of 5
	// slice := []int{10, 20, 30, 40, 50}
	// //create a new slice of a length 2 and capacity of 4 elements
	// newSlice := slice[1:3]
	// fmt.Println(newSlice)
	// fmt.Println(cap(newSlice))

	// //Add new element to the slice using append:
	// newSlice = append(newSlice, 100)
	// fmt.Println(newSlice)

	//Performing a three-index slice
	//Slice the third element and restrict the capacity
	//Containa a length of 1 element and capacity of 2 elements
	// source := []string{"Apple", "Mango", "Orange", "Banana"}
	// slice2 := source[2:3:4]
	//For slice[i:j:k]
	// Length j - i 3-2 =1
	// Capacity k-i 4-2 =2
	// fmt.Println(slice2)
	// fmt.Println(cap(slice2))

	// //Append a to a slice from another a slice
	// s1 := []int{1, 2}
	// s2 := []int{3, 4}
	// s3 := append(s1, s2...)
	// fmt.Println(s3)

	// for index, value := range s3 {
	// 	fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &s3[index])
	// }

	// // Iterate over each element starting at elements 3
	// for index := 2; index < len(s3); index++ {
	// 	fmt.Printf("Index: %d Value: %d \n", index, s3[index])
	// }

	//declaring a multi-dimensional slice
	slice := [][]int{{10}, {100, 200}}
	fmt.Println(slice)
	slice[0] = append(slice[0], 20)
	fmt.Println(slice)

	//PASSING SLICE BETWEEN FUNCTIONs
	//Allocate a slice of 1 million integers
	// slice2 := make([]int, 1e6)
	//passing the slice to the function foo
	// slice = foo(slice2)

	//function foo accepts a slice of integers and returns the slice back
	//Note That we dont have to use pointers here
	// func foo(slice2 []int) []int {
	// 	return slice2
	// }

}
