package user

import _entities "gokomodo-api/entities"

type UserRepositoryInterface interface {
	CreateSeller(request _entities.Seller) (_entities.Seller, error)
	CreateBuyer(request _entities.Buyer) (_entities.Buyer, error)
}
