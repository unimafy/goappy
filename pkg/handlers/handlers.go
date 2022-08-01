package handlers

import (
	"myapp/pkg/renderers"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "/home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "/about.html")
}
