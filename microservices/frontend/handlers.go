package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func (f *frontendServer) getCartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cart handler reached")

	// TODO: IMPLEMENT this which will fetch items from shopping cart for given user
	/*cartProducts, err := getCartProductsForUser("user id")
	if err != nil {
		fmt.Println(err)
		// todo: display proper page
	}*/

	err := templates.ExecuteTemplate(w, "cart", nil)
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
	}
}

func (f *frontendServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Signin handler reached")

	err := templates.ExecuteTemplate(w, "login", nil)
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
	}
}

func (f *frontendServer) postLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Signin post handler reached")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}
	fmt.Println("email:", r.Form["email"])
	fmt.Println("password:", r.Form["password"])

	// redirect the user to the Home Page after he logged in
	//http.RedirectHandler("/", 200)
	// TODO: DOES NOW WORK, IT WILL REDIRECT HOME, BUT NO PRODUCTS WILL BE LOADED

	err = templates.ExecuteTemplate(w, "login_result", "successfull")
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
	}
}

func (f frontendServer) productDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Product details handler reached")

	var id = mux.Vars(r)["id"]

	product, err := getProductById(id)
	if err != nil {
		fmt.Println(err)
		// todo: assign empty array of skateboard and sed it to frontend
	}
	fmt.Println(product)

	err = templates.ExecuteTemplate(w, "product", product)
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
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

func getProductById(id string) (Product, error) {
	// TODO: Replace url: with Kubernetes service name when deploying in K8s env
	var product Product

	c := http.Client{Timeout: time.Duration(3) * time.Second}

	url := fmt.Sprintf("http://localhost:3030/product/%v", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return product, err
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return product, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return product, err
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		return product, err
	}

	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

	return product, nil

}
