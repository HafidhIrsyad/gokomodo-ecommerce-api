package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		authUseCase := NewAuthService(mockAuthRepository{})
		data, err := authUseCase.LoginSeller("", "")
		assert.Nil(t, err)
		assert.Equal(t, "odi@mail.com", data)
	})

	t.Run("TestGetLoginSellerError", func(t *testing.T) {
		authUseCase := NewAuthService(mockAuthRepository{})
		data, err := authUseCase.LoginSeller("", "")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})

	t.Run("TestLoginBuyerSuccess", func(t *testing.T) {
		authUseCase := NewAuthService(mockAuthRepository{})
		data, err := authUseCase.LoginBuyer("", "")
		assert.Nil(t, err)
		assert.Equal(t, "hafidh@mail.com", data)
	})

	t.Run("TestGetLoginBuyerError", func(t *testing.T) {
		authUseCase := NewAuthService(mockAuthRepository{})
		data, err := authUseCase.LoginBuyer("", "")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

// === mock success ===
type mockAuthRepository struct{}

func (m mockAuthRepository) LoginBuyer(email string, password string) (string, error) {
	//TODO implement me
	return "hafidh@mail.com", nil
}

func (m mockAuthRepository) LoginSeller(email string, password string) (string, error) {
	return "hafidh@mail.com", nil
}

// === mock error ===

type mockAuthRepositoryError struct{}

func (m mockAuthRepositoryError) LoginSeller(email string, password string) (string, error) {
	return "", fmt.Errorf("error")
}

func (m mockAuthRepositoryError) LoginBuyer(email string, password string) (string, error) {
	return "", fmt.Errorf("error")
}
