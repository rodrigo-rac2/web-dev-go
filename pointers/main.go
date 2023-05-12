package main

import "log"

func main() {
	var myString string = "Green"
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after function call, myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	log.Println("s is set to", s, ", which is the pointer")
	*s = newValue
}
