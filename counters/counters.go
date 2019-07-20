package counters

//unexported type:
type alertCounter int

//New creates and returns values of the unexpected:
func New(value int) alertCounter {
	return alertCounter(value)
}

type user struct {
	Name  string
	Email string
}
type Admin struct {
	user   //unexported
	Rights int
}
