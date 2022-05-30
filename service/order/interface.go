package order

import _entities "gokomodo-api/entities"

type OrderServiceInterface interface {
	CreateOrder(order _entities.Order) (_entities.Order, error)
	GetAllSellerOrder(id int) ([]_entities.Order, error)
	GetAllBuyerOrder(id int) ([]_entities.Order, error)
}
