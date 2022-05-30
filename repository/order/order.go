package order

import (
	_entities "gokomodo-api/entities"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (or *OrderRepository) GetAllSellerOrder(id int) ([]_entities.Order, error) {
	var order []_entities.Order

	tx := or.DB.Where("seller_id = ?", id).Find(&order)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return order, nil
}

func (or *OrderRepository) GetAllBuyerOrder(id int) ([]_entities.Order, error) {
	var order []_entities.Order

	tx := or.DB.Where("buyer_id = ?", id).Find(&order)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return order, nil
}

func (or *OrderRepository) CreateOrder(request _entities.Order) (_entities.Order, error) {
	yx := or.DB.Save(&request)

	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}
