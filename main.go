package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Application is running on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
