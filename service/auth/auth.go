package auth

import (
	"gokomodo-api/repository/auth"
)

type AuthService struct {
	authRepository auth.AuthRepositoryInterface
}

func NewAuthService(authRepo auth.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}

func (as *AuthService) LoginSeller(email string, password string) (string, error) {
	token, err := as.authRepository.LoginSeller(email, password)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) LoginBuyer(email string, password string) (string, error) {
	token, err := as.authRepository.LoginBuyer(email, password)

	if err != nil {
		return "", err
	}

	return token, nil
}
