package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

var templates = template.Must(template.New("").ParseGlob("templates/*.html"))

func (f *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler reached")

	getAllProducts("site")
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		fmt.Println("Error rendering page")
		// todo: implement a renderHttpErrorPage
	}
}

// private helper functions used to fetch data form the other microservices

func getAllProducts(endpoint string) {
	// TODO: Replace url: with Kubernetes service name when deploying in K8s env

	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:3030/product", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	// TODO: Map response body to an array of type []product
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)
}
