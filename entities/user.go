package entities

import (
	"gorm.io/gorm"
)

type Buyer struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	BuyerAddress string `json:"buyer_address" form:"buyer_address"`
}

type Seller struct {
	gorm.Model
	Name          string `json:"name" form:"name"`
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	PickupAddress string `json:"pickup_address" form:"pickup_address"`
}
