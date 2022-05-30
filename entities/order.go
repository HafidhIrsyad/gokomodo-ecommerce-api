package entities

import "gorm.io/gorm"

type Order struct {
	*gorm.Model
	SellerID                   int     `json:"seller_id"`
	BuyerID                    int     `json:"buyer_id"`
	DeliverySourceAddress      string  `json:"delivery_source_address"`
	DeliveryDestinationAddress string  `json:"delivery_destination_address"`
	Items                      string  `json:"items"`
	Quantity                   int     `json:"quantity"`
	Price                      int     `json:"price"`
	TotalPrice                 int     `json:"total_price"`
	Status                     string  `json:"status"`
	Buyer                      Buyer   `json:"buyer" gorm:"foreignKey:BuyerID;references:ID"`
	Seller                     Seller  `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
	Product                    Product `json:"product" gorm:"foreignKey:Items;references:ID"`
}
