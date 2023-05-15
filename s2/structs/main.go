package main

import "log"

type User struct {
	id        int
	firstName string
	lastName  string
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetId(newId int) {
	u.id = newId
}

func (u *User) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) SetName(newFirstName string, newLastName string) {
	u.firstName = newFirstName
	u.lastName = newLastName
}

func main() {
	var user User = User{1, "John", "Smith"}
	log.Println("The user's full name is", user.GetFullName())
	user.SetName("Jane", "Doe")
	log.Println("The user's full name was updated to", user.GetFullName())

	// no privacy within the package
	user.firstName = "John"
	user.lastName = "Smith"
	log.Println("The user's full name was updated to", user.GetFullName())
}
