package main

import "fmt"

//User defined types:
type user struct {
	name       string
	email      string
	ext        int
	priveleged bool
}

//Declare a variable of type user
var bill user

func main() {
	//declaring a variable and initializing the fields:
	Lisa := user{
		name:       "Lisa",
		email:      "lisa@gmail.com",
		ext:        123,
		priveleged: false,
	}

	fmt.Println(Lisa)

	//An admin is a user with privelges
	type admin struct {
		person user
		level  string
	}

	fred := admin{
		person: user{
			name:       "Fred",
			email:      "fred@admin.com",
			ext:        123,
			priveleged: true,
		},
		level: "super",
	}
	fmt.Println(fred)

	// hello()

	var dur Duration
	dur = 1000
	fmt.Println(dur)

}

// func hello(p ...int) {
// 	fmt.Printf("%T\n", p) //this outputs a slice of int
// 	if p == nil {
// 		fmt.Println("nil")
// 		return
// 	}
// 	fmt.Println("not nil")
// }

type Duration int64
