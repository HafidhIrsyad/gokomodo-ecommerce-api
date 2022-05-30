package user

import (
	"gokomodo-api/entities"
	_userRepository "gokomodo-api/repository/user"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserService(userRepo _userRepository.UserRepositoryInterface) UserServicedInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (us UserService) CreateSeller(request entities.Seller) (entities.Seller, error) {
	//TODO implement me
	users, err := us.userRepository.CreateSeller(request)

	return users, err
}

func (us UserService) CreateBuyer(request entities.Buyer) (entities.Buyer, error) {
	//TODO implement me
	users, err := us.userRepository.CreateBuyer(request)

	return users, err
}
