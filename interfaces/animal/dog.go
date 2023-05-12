package animal

// Dog defines the dog type
type Dog struct {
	Name  string
	Breed string
}

// Says has a receiver of type *Dog, so it satisfies part of the interface requirements for Animal
// for the Dog type
func (d *Dog) Says() string {
	return "Woof"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Dog type
func (d *Dog) NumberOfLegs() int {
	return 4
}
