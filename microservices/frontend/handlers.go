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
	fmt.Println("I AM BACK ON HOME HANDLER")
	fmt.Println("CALLING PRODUCT 1")
	//fmt.Println(products[1]) Error was here
	fmt.Println("AND HERE IS THE ERROR. TOLD YOU")

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

	token, err := loginUser(r.Form["email"][0], r.Form["password"][0])
	if err != nil {
		templates.ExecuteTemplate(w, "login_result", "unsuccessful, please try again later.")
	}

	// redirect the user to the Home Page after he logged in
	//http.RedirectHandler("/", 200)
	// TODO: DOES NOW WORK, IT WILL REDIRECT HOME, BUT NO PRODUCTS WILL BE LOADED

	// Save token in cookie
	expiration := time.Now().Add(10 * time.Minute)
	cookie := http.Cookie{Name: "skateshop_login", Value: token, Expires: expiration}
	http.SetCookie(w, &cookie)

	// Read cookie
	cookieM, _ := r.Cookie("skateshop_login")
	fmt.Println(fmt.Sprintf("I retrieved the following cookie: %v", cookieM))

	err = templates.ExecuteTemplate(w, "login_result", token)
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

// Add item to cart requires us to be authorized
func (f frontendServer) addItemToCartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Product details handler reached")

	var id = mux.Vars(r)["id"]
	fmt.Println(fmt.Sprintf("Add to cart item with id: %v", id))

	// Read cookie
	cookieM, err := r.Cookie("skateshop_login")
	if err != nil {
		fmt.Println("No cookie")
		templates.ExecuteTemplate(w, "error", "You have no cookie")
		// TODO: Throw 404 Unauthorized
	}
	fmt.Println(fmt.Sprintf("Auth cookie: %v", cookieM))

	// Just display the cart for now, later we do no redirect actually
	err = templates.ExecuteTemplate(w, "cart", nil)
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
	}

	// Bug: Both templates get executed in case there is no cookie.
}

// private helper functions used to fetch data form the other microservices

func getAllProducts(endpoint string) ([]Product, error) {
	// TODO: Replace url: with Kubernetes service name when deploying in K8s env
	var products []Product

	fmt.Printf("getAllProducts reached fine")

	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, err := http.NewRequest("GET", "http://catalog.skateshop.svc.cluster.local:3030/product", nil)
	if err != nil {
		fmt.Println("I GET ERROR WHILE CALLING K8S SERVICE")
		fmt.Printf("error %s", err)
		return products, err
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)

	fmt.Printf("I DID THE REQUEST I AM WAITING FOR ERRORS")

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

	fmt.Printf("NO ERROR???")
	return products, nil

}

func getProductById(id string) (Product, error) {
	
	var product Product

	c := http.Client{Timeout: time.Duration(3) * time.Second}

	url := fmt.Sprintf("http://catalog.skateshop.svc.cluster.local:3030/product/%v", id)
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

func loginUser(username string, password string) (string, error) {
	// TODO: Replace url: with Kubernetes service name when deploying in K8s env
	var token string

	c := http.Client{Timeout: time.Duration(3) * time.Second}

	url := "http://localhost:8070/login"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return token, err
	}
	req.Header.Add("Accept", `text/plain`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return token, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return token, err
	}

	token = string(body)
	fmt.Printf(fmt.Sprintf("Got token: %v", token))

	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

	return token, nil

}
