package main

import "fmt"

type UserCart struct {
	UserID string
	Items  []Product
	Total  int
}

type Product struct {
	ProductID string
	Name      string `json:"name"`
	Price     int    `json:"price"`
}

var CartInMemoryDB = []UserCart{
	{
		UserID: "34",
		Items: []Product{
			{
				ProductID: "1",
				Name:      "Skate Cliche",
				Price:     140,
			},
		},
		Total: 140,
	},
	{
		UserID: "1",
		Items: []Product{
			{
				ProductID: "1",
				Name:      "Skate Cliche",
				Price:     140,
			},
		},
		Total: 140,
	},
}

func CalculateTotalForUserID() {
	// TODO: Implement in a next version of the application
}

func CheckoutOrder(UserID string) {

	ok := false
	pos := 0

	for i := 0; i < len(CartInMemoryDB); i++ {
		if CartInMemoryDB[i].UserID == UserID {
			fmt.Println(fmt.Sprintf("Found my user at index %v", i))
			ok = true
			pos = i
		}
	}

	if ok == true {
		CartInMemoryDB = append(CartInMemoryDB[:pos], CartInMemoryDB[pos+1:]...)
	}
}

func main() {
	fmt.Println("Before removing")

	fmt.Println(fmt.Sprintf("%+v", CartInMemoryDB))

	CheckoutOrder("1")
	fmt.Println("After removing")

	fmt.Println(fmt.Sprintf("%+v", CartInMemoryDB))
}
