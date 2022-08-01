package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application is running on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
