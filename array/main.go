package main

import "fmt"

func main() {
	// theArray := [4]int{1, 2, 4, 5}
	theArray := [...]int{1, 2, 4, 5}

	// fmt.Println(theArray)
	for index, value := range theArray {
		if index == 2 {
			continue
		}
		fmt.Println(value) //this will print only three elements in the array
	}

	//Initialize index 1 and 2 with specific values
	//The rest of the elements contain their zero value.
	theSecondArray := [5]int{1: 10, 2: 20}

	fmt.Println(theSecondArray)

	//Accessing array pointer elements:
	//Declare an integer pointer array of five elements
	//Initialize index 0 and 1 of the array with integer pointers.

	array1 := [5]*int{0: new(int), 1: new(int)}
	//Asign values to index 0 and 1
	*array1[0] = 10
	*array1[1] = 20

	fmt.Println(array1)
	// fmt.Println(array)

	var array2 [5]*int

	array2 = array1

	fmt.Println(array2)

	//Multi dimensional array
	// var multiArray = [4][2]int

	multiArray := [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}

	fmt.Println(multiArray)

	//Decalre and initialize the 1 and 3 elements outer array
	multiArray2 := [4][2]int{1: {20, 21}, 3: {40, 41}}

	fmt.Println(multiArray2)

	//Declare and initialize elements of the outer and inner arrays:
	multiArray3 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Println(multiArray3)

	//Set values for each individual elements:
	var multiArray4 [4][2]int

	multiArray4[0][0] = 20
	multiArray4[0][1] = 30
	multiArray4[1][0] = 40
	multiArray4[1][1] = 50

	fmt.Println(multiArray4)

	//Assigning multi dimensional arrays by index:

	var singleFromMulti [2]int = multiArray4[1]

	fmt.Println(singleFromMulti)

	//Passing arrays between functions
	//Passing a large array by pointer between functions
	//Allocate an array of 8 megabytes
	var arrayFunc [1e6]int
	//pass the address of the array to the function foo
	foo(&arrayFunc)

	//Function foo accepts a pointer to an array of one million integers

}

func foo(arrayFunc *[1e6]int) {
}
