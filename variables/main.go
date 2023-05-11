package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var sayGoodbye string = "Goodbye"
	var numberA int = 4
	fmt.Println(sayGoodbye)

	numberB := 5
	fmt.Println("The sum of the numbers is", numberA+numberB)

	something, somethingElse := printSomething()
	fmt.Println(something, somethingElse)
}

func printSomething() (string, string) {
	return "Something", "Something else"
}
