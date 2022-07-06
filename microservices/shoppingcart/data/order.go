package data

import "fmt"

type UserCart struct {
	UserID string
	Items  []Product
}

type Product struct {
	ProductID string `json:"id"`
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
	},
}

func GetItemsForUserID() []Product {
	for i := 0; i < len(CartInMemoryDB); i++ {
		if CartInMemoryDB[i].UserID == "1" {
			fmt.Println("Found items for user ID")
			return CartInMemoryDB[i].Items
		}
	}

	fmt.Printf("No items found for user ID")
	return []Product{}
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

func AddItemToCart() {
	fmt.Println("Item added to cart")
}
