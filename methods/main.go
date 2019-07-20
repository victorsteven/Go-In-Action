package main

import "fmt"

type user struct {
	name  string
	email string
}

//notify implements a method with a value receiver
func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

//changeEmail implements a method with a pointer receiver:
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	//Values of type user can be used to call methods
	//declared with a value receiver
	// steven := user{
	// 	name: "Steven", email: "steven@gmail.com",
	// }
	steven := user{"Steven", "steven@gmail.com"}
	steven.notify()

	//Pointers of type user can also be used to call methods
	//declare with a value receiver

	bill := &user{"Bill", "bill@gmail.com"}

	bill.notify() //what happens underneatha; (*bill).notify()

	//Values of type user can also be used to call methods
	//declared with a pointer receiver
	steven.changeEmail("secondsteven@gmail.com")
	steven.notify() //what happens underneath: (&steven).notify()

	//Pointers of type user can be used to call methods
	//declared with a pointer receiver

	bill.changeEmail("billy@gmail.com")
	bill.notify()

	// Determining whether to use a value or pointer receiver can sometimes be confus- ing.

}
