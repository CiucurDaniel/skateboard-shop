package data

type ShoppingCart struct {
	UserID string
	Items  []Product
	Total  int
}

type Product struct {
	ProductID string
	Name      string `json:"name"`
	Price     int    `json:"price"`
}

var CartInMemoryDB = []ShoppingCart{
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

func (c ShoppingCart) CalculateTotalForUserID() {
	// TODO: Implement
}

func (c ShoppingCart) PlaceOrder() {
	// TODO: Implement
	// Info: This will just delete the items for the user for now
}
