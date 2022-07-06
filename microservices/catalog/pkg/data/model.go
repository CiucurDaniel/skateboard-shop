package data

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

var DemoProducts = []Product{p1, p2, p3, p4, p5}

var p1 = Product{
	Name:        "Cliche Pro",
	Price:       140,
	Quantity:    4,
	ImageUrl:    "skate_cliche.PNG",
	Description: "A nice skateboard",
}

var p2 = Product{
	Name:        "Baker",
	Price:       200,
	Quantity:    4,
	ImageUrl:    "skate_baker.PNG",
	Description: "A nice skateboard",
}
var p3 = Product{
	Name:        "Toy Machine",
	Price:       240,
	Quantity:    4,
	ImageUrl:    "skate_toy_machine.PNG",
	Description: "A nice skateboard",
}

var p4 = Product{
	Name:        "Element",
	Price:       130,
	Quantity:    4,
	ImageUrl:    "skate_element.PNG",
	Description: "A nice skateboard",
}

var p5 = Product{
	Name:        "Jart",
	Price:       140,
	Quantity:    4,
	ImageUrl:    "skate_jart.PNG",
	Description: "A nice skateboard",
}
