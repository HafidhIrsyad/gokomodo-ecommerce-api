package product

import _entities "gokomodo-api/entities"

type ProductServiceInterface interface {
	CreateProduct(product _entities.Product) (_entities.Product, error)
	GetAllProduct() ([]_entities.Product, error)
}
