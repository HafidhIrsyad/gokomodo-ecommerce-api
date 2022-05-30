package handler

import (
	"github.com/labstack/echo/v4"
	_middlewares "gokomodo-api/delivery/middlerwares"
	"gokomodo-api/delivery/response"
	_entities "gokomodo-api/entities"
	"gokomodo-api/service/order"
	"net/http"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func NewOrderHandler(orderService order.OrderServiceInterface) OrderHandler {
	return OrderHandler{orderService}
}

func (oh *OrderHandler) GetOrderSellerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		order, err := oh.orderService.GetAllSellerOrder(idToken)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("failed get all order", http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all order", http.StatusOK, order))
	}
}

func (oh *OrderHandler) GetOrderBuyerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		orderBuyer, err := oh.orderService.GetAllBuyerOrder(idToken)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("failed get all order", http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all order", http.StatusOK, orderBuyer))
	}
}

func (oh *OrderHandler) CreateBookingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var order _entities.Order

		idToken, errToken := _middlewares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, response.ResponseFailed("Unauthorized", http.StatusUnauthorized))
		}

		order.BuyerID = idToken

		errBind := c.Bind(&order)

		if errBind != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed(errBind.Error(), http.StatusBadRequest))
		}

		if order.TotalPrice <= 0 {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Total price must be greater than 0", http.StatusBadRequest))
		}

		newOrderBuyer, errCreate := oh.orderService.CreateOrder(order)

		if errCreate != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("created booking failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("created booking successfully", http.StatusOK, newOrderBuyer))
	}
}
