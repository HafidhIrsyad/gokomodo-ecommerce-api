package handler

import (
	"github.com/labstack/echo/v4"
	"gokomodo-api/delivery/response"
	_entities "gokomodo-api/entities"
	userService "gokomodo-api/service/user"
	"net/http"
)

type UserHandler struct {
	userService userService.UserServicedInterface
}

func NewUserHandler(us userService.UserServicedInterface) UserHandler {
	return UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) CreateSellerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var param _entities.Seller

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error binding data", http.StatusBadRequest))
		}
		if _, ok := response.ValidMailAddress(param.Email); !ok {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error to validate email", http.StatusBadRequest))
		}
		_, err := uh.userService.CreateSeller(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Register failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("Successfully registered", http.StatusOK))
	}
}

func (uh *UserHandler) CreateBuyerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var param _entities.Buyer

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error binding data", http.StatusBadRequest))
		}
		if _, ok := response.ValidMailAddress(param.Email); !ok {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error to validate email", http.StatusBadRequest))
		}
		_, err := uh.userService.CreateBuyer(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Register failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("Successfully registered", http.StatusOK))
	}
}
