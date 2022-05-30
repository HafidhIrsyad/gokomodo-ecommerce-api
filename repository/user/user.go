package user

import (
	"errors"
	password2 "gokomodo-api/delivery/password"
	_entities "gokomodo-api/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateSeller(request _entities.Seller) (_entities.Seller, error) {
	if request.Name == "" || request.Name == " " {
		return request, errors.New("can't be empty")
	}
	if request.Email == "" || request.Email == " " {
		return request, errors.New("can't be empty")
	}
	if request.Password == "" || request.Password == " " {
		return request, errors.New("can't be empty")
	}
	if request.PickupAddress == "" || request.PickupAddress == " " {
		return request, errors.New("can't be empty")
	}

	password, _ := password2.HashPassword(request.Password)
	request.Password = password

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}

func (ur *UserRepository) CreateBuyer(request _entities.Buyer) (_entities.Buyer, error) {
	if request.Name == "" || request.Name == " " {
		return request, errors.New("can't be empty")
	}
	if request.Email == "" || request.Email == " " {
		return request, errors.New("can't be empty")
	}
	if request.Password == "" || request.Password == " " {
		return request, errors.New("can't be empty")
	}
	if request.BuyerAddress == "" || request.BuyerAddress == " " {
		return request, errors.New("can't be empty")
	}

	password, _ := password2.HashPassword(request.Password)
	request.Password = password

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}
