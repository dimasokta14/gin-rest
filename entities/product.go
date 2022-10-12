package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `json:"title"`
	Desription string `json:"description"`
	SKU        string `json:"sku"`
	Stock      int    `json:"stock"`
	ImageUrl   string `json:"image_url"`
}

type AddProductRequestBody struct {
	Title      string `json:"title"`
	Desription string `json:"description"`
	SKU        string `json:"sku"`
	Stock      int    `json:"stock"`
	ImageUrl   string `json:"image_url"`
}
