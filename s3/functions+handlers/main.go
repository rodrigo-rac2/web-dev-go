package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":9600"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValues adds two ints x and y, and returns the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100
	var y float32 = 0
	//var y float32 = 20
	result, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Println("[ERR! DIVIDE] " + err.Error())
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, result))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0.0 {
		err := errors.New("Cannot divide by zero")
		return 0.0, err
	}
	return x / y, nil
}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s...", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
	fmt.Println("Shutting down server...")
	fmt.Println("Done!")
}
