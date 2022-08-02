package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Application is running on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
