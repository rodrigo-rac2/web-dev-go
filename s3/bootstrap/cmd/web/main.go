package main

import (
	"fmt"
	"github.com/rodrigo-rac2/web-dev-go/s3/bootstrap/pkg/handlers"
	"net/http"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting server on http://localhost" + portNumber + "/ ...")
	_ = http.ListenAndServe(portNumber, nil)
}
