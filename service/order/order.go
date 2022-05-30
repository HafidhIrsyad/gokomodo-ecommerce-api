package order

import (
	_entities "gokomodo-api/entities"
	"gokomodo-api/repository/order"
)

type OrderService struct {
	orderRepo order.OrderRepositoryInterface
}

func NewOrderService(orderRepo order.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (os *OrderService) GetAllSellerOrder(id int) ([]_entities.Order, error) {
	order, err := os.orderRepo.GetAllSellerOrder(id)

	return order, err
}

func (os *OrderService) GetAllBuyerOrder(id int) ([]_entities.Order, error) {
	order, err := os.orderRepo.GetAllBuyerOrder(id)

	return order, err
}

func (os *OrderService) CreateOrder(request _entities.Order) (_entities.Order, error) {
	order, err := os.orderRepo.CreateOrder(request)

	return order, err
}
