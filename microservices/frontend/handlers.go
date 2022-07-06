package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	//fmt.Println(products[1]) Error was here

	err = templates.ExecuteTemplate(w, "home", products)
	if err != nil {
		fmt.Println("Error rendering page")
		// todo: implement a renderHttpErrorPage
	}
}

func (f *frontendServer) getCartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cart handler reached")

	products := getCartItems("1")

	err := templates.ExecuteTemplate(w, "cart", products)
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

	didRenderOneTemplate := false

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}
	fmt.Println("id:", r.Form["id"])
	fmt.Println("name:", r.Form["name"])
	fmt.Println("price:", r.Form["price"])

	// Read cookie
	cookieM, err := r.Cookie("skateshop_login")
	if err != nil {
		fmt.Println("No cookie")
		templates.ExecuteTemplate(w, "error", "You have no cookie. You have to login first in order to access your cart")
		didRenderOneTemplate = true
		// TODO: Throw 404 Unauthorized
	}
	fmt.Println(fmt.Sprintf("Auth cookie: %v", cookieM))

	priceint, _ := strconv.Atoi(r.Form["price"][0])

	// Here make 2 requests to cart microservice
	addProductToCart(r.Form["name"][0], priceint)
	products := getCartItems("1")

	if didRenderOneTemplate == false {
		// Just display the cart for now, later we do no redirect actually
		err = templates.ExecuteTemplate(w, "cart", products)
		if err != nil {
			templates.ExecuteTemplate(w, "error", err.Error())
		}
	}

	// Bug: Both templates get executed in case there is no cookie.
}

func (f *frontendServer) postCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checkout post handler reached")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}
	fmt.Println("card:", r.Form["card"])
	fmt.Println("cvc:", r.Form["cvc"])
	fmt.Println("address:", r.Form["address"])

	// TODO: Call shopping cart with user id from cookie the shopping cart will simply accept the payment meaning it will delete all items from the cart
	err = checkoutForUserId("1")
	if err != nil {
		fmt.Println("Could not successfully checkout order")
		templates.ExecuteTemplate(w, "checkout_result", "NOT successful")
	}

	err = templates.ExecuteTemplate(w, "checkout_result", "SUCCESSFUL")
	if err != nil {
		templates.ExecuteTemplate(w, "error", err.Error())
	}
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

	url := "http://auth.skateshop.svc.cluster.local:8070/login"
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

func checkoutForUserId(id string) error {
	c := http.Client{Timeout: time.Duration(3) * time.Second}

	url := "http://shoppingcart.skateshop.svc.cluster.local:8060/checkout"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return err
	}
	req.Header.Add("Accept", `text/plain`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return err
	}

	fmt.Println(fmt.Sprintf("Got response: %v", resp.Status))
	return nil
}

func addProductToCart(name string, price int) {
	c := http.Client{Timeout: time.Duration(3) * time.Second}

	// assume user id = 1 added the item

	p := Product{
		Name:  name,
		Price: price,
	}

	pjson, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Could not conv to json your cart item")
	}

	url := "http://shoppingcart.skateshop.svc.cluster.local:8060/additem"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pjson))
	if err != nil {
		fmt.Printf("error %s", err)
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	fmt.Println(fmt.Sprintf("Got response: %v", resp.Status))
}

func getCartItems(userid string) []Product {
	fmt.Println("HERE TAKE YOUR ITEMS")

	var products []Product

	c := http.Client{Timeout: time.Duration(3) * time.Second}

	// assume user id = 1 added the item

	url := "http://shoppingcart.skateshop.svc.cluster.local:8060/cart"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Printf("Error converting body response to json")
	}

	return products
}
