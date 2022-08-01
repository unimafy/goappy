package renderers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(
		"./templates/"+tmpl,
		"./templates/base.layout.html",
	)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the template", err)
		return
	}
}
