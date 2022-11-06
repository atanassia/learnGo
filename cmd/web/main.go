package main

import (
	"fmt"
	"net/http"
	"github.com/atanassia/go-course/pkg/handlers"
)


const portNumber = ":8000"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}