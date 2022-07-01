package main

import (
	"encoding/json"
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

	products, err := getAllProducts("site")
	if err != nil {
		fmt.Println(err)
		// todo: assign empty array of skateboard and sed it to frontend
	}
	fmt.Println(products[1])

	err = templates.ExecuteTemplate(w, "home", products)
	if err != nil {
		fmt.Println("Error rendering page")
		// todo: implement a renderHttpErrorPage
	}
}

// private helper functions used to fetch data form the other microservices

func getAllProducts(endpoint string) ([]Product, error) {
	// TODO: Replace url: with Kubernetes service name when deploying in K8s env
	var products []Product

	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:3030/product", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return products, err
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return products, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return products, err
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		return products, err
	}

	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

	return products, nil

}
