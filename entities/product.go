package entities

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	Seller      int    `json:"seller"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Sellers     Seller `json:"product" gorm:"foreignKey:Seller;references:ID"`
}
