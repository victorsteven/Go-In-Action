package main

import "fmt"

// notifier is an inteface that defined notification type behavior
type notifier interface {
	notify()
}

// user definees a user in the program
type user struct {
	name  string
	email string
}
type admin struct {
	name  string
	email string
}

//using type embedding:
type agent struct {
	user  //type embedding
	level bool
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// notify implements a method with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending message to %s<%s>\n", u.name, u.email)
}

func main() {
	// Create a value of type User and send a notification
	u := user{"adam", "adam@gmail.com"}
	lisa := admin{"Lisa", "lisa@gmail.com"}

	ag := agent{
		user: user{
			name:  "Secret Agent",
			email: "secret@gmail.com",
		},
		level: true,
	}
	// ag.notify() //using type promotion
	//note that the agent type never implemented the notifier interface, it could use it because of the user type that implements the interface

	sendNotification(&ag)

	sendNotification(&lisa)

	//only pointers of type user can be passed to the sendNotification function, since a pointer receiver was used to implement the interface

	sendNotification(&u)

	//send notification Accept a valur that implements the notifier interface and sends notifications

}

func sendNotification(n notifier) {
	n.notify()
}
