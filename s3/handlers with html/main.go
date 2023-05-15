package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "about.page.tmpl")
}

func renderHtml(w http.ResponseWriter, htmlTemplate string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + htmlTemplate)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting server on http://localhost" + portNumber + "/ ...")
	_ = http.ListenAndServe(portNumber, nil)
}
