package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.New("").ParseGlob("templates/*.html"))

func (f *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler reached")

	// Note: templates must not be loaded for every request, change in future
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		fmt.Println("Error rendering page")
		// todo: implement a renderHttpErrorPage
	}
}
