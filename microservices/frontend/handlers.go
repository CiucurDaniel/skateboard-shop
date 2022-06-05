package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (f *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Home handler reached")

	// Note: templates must not be loaded for every request, change in future
	err := template.Must(template.New("").ParseGlob("templates/*.html")).ExecuteTemplate(w,"home", nil)
	if err != nil {
		fmt.Println("Error rendering page")
		// todo: implement a renderHttpErrorPage
	}
}
