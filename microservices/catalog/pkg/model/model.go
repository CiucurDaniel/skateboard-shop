package model

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	ImageUrl string `json:"image_url"`
	Description string `json:"description"`
}