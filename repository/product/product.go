package product

import (
	_entities "gokomodo-api/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (pr *ProductRepository) GetAllProduct() ([]_entities.Product, error) {
	var product []_entities.Product

	tx := pr.DB.Find(&product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func (pr *ProductRepository) CreateProduct(request _entities.Product) (_entities.Product, error) {
	yx := pr.DB.Save(&request)

	if yx.Error != nil {
		return request, yx.Error
	}

	return request, nil
}