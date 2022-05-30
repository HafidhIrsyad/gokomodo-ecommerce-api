package product

import (
	_entities "gokomodo-api/entities"
	"gokomodo-api/repository/product"
)

type ProductService struct {
	productRepository product.ProductRepositoryInterface
}

func NewProductService(productRepo product.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (p ProductService) CreateProduct(product _entities.Product) (_entities.Product, error) {
	//TODO implement me
	product, err := p.productRepository.CreateProduct(product)

	return product, err
}

func (p ProductService) GetAllProduct() ([]_entities.Product, error) {
	//TODO implement me
	product, err := p.productRepository.GetAllProduct()

	return product, err
}
