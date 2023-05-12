package animal

// Gorilla defines the Gorilla type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Says has a receiver of type *Gorilla, so it satisfies part of the interface requirements for Animal
// for the Gorilla type
func (d *Gorilla) Says() string {
	return "Ugh"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Gorilla type
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
