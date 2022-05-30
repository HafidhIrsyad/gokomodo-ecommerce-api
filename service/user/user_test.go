package user

import (
	"github.com/stretchr/testify/assert"
	_entities "gokomodo-api/entities"
	"testing"
)

func TestCreateSeller(t *testing.T) {
	t.Run("TestCreateSellerSuccess", func(t *testing.T) {
		userUseCase := NewUserService(mockUserRepository{})
		data, err := userUseCase.CreateBuyer(_entities.Buyer{Name: "haf"})
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.Name)
	})

	t.Run("TestCreateSellerSuccess", func(t *testing.T) {
		userUseCase := NewUserService(mockUserRepositoryError{})
		data, err := userUseCase.CreateBuyer(_entities.Buyer{Name: "haf"})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Name)
		assert.Nil(t, nil, err)
	})
}

func TestCreateBuyer(t *testing.T) {
	t.Run("TestCreateBuyerSuccess", func(t *testing.T) {
		userUseCase := NewUserService(mockUserRepository{})
		data, err := userUseCase.CreateBuyer(_entities.Buyer{Name: "haf"})
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.Name)
	})

	t.Run("TestCreateBuyerSuccess", func(t *testing.T) {
		userUseCase := NewUserService(mockUserRepositoryError{})
		data, err := userUseCase.CreateBuyer(_entities.Buyer{Name: "haf"})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Name)
		assert.Nil(t, nil, err)
	})
}

type mockUserRepository struct{}

func (m mockUserRepository) CreateSeller(request _entities.Seller) (_entities.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockUserRepository) CreateBuyer(request _entities.Buyer) (_entities.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) CreateSeller(request _entities.Seller) (_entities.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockUserRepositoryError) CreateBuyer(request _entities.Buyer) (_entities.Buyer, error) {
	//TODO implement me
	panic("implement me")
}
