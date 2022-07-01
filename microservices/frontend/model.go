package main

type Product struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}
