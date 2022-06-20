package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}
