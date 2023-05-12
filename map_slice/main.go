package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	// maps
	users := make(map[string]User)
	users["Smith"] = User{"John", "Smith"}
	users["Doe"] = User{"Jane", "Doe"}
	me := User{
		FirstName: "Rodrigo",
		LastName:  "Costa",
	}
	users["Costa"] = me
	log.Println(users)
	selectedUser := users["Costa"]
	log.Println("The user's full name is", selectedUser.GetFullName())

	// slices
	var intSlice []int
	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 3)
	log.Println("The first element is", intSlice[0])
	sort.Ints(intSlice)
	log.Println("The ordered slice is", intSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers[0])

	//print first two elements of the numbers slice
	log.Println(numbers[:2])

	//print last two elements of the numbers slice
	log.Println(numbers[len(numbers)-2:])

	//print the last element of the numbers slice
	log.Println(numbers[len(numbers)-1])

	//print the second to the fifth elements of the numbers slice
	log.Println(numbers[1:5])

}
