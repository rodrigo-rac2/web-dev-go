package main

import (
	"fmt"
	animal2 "github.com/rodrigo-rac2/web-dev-go/s2/interfaces/animal"
)

func main() {
	dog := animal2.Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	// We can pass dog to PrintInfo(), since the Dog type implements the Animal interface by having all of the
	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
	// https://tour.golang.org/methods/8 for more details).
	PrintInfo(&dog)

	gorilla := animal2.Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	// We can also pass gorilla to PrintInfo(), since the Gorilla type implements the Animal interface by having all of the
	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
	// https://tour.golang.org/methods/8 for more details).
	PrintInfo(&gorilla)
}

func PrintInfo(a animal2.Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
