package handler

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gokomodo-api/delivery/response"
	"gokomodo-api/entities"
	"gokomodo-api/service/auth"
	"net/http"
)

type AuthHandler struct {
	authUseCase auth.AuthServiceInterface
}

func NewAuthHandler(auth auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: auth,
	}
}

func (ah *AuthHandler) LoginSellerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login entities.Seller
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error to bind data", http.StatusBadRequest))
		}
		token, errorLogin := ah.authUseCase.LoginSeller(login.Email, login.Password)

		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed(fmt.Sprintf("%v", errorLogin), http.StatusBadRequest))
		}
		extract, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("S3CR3T"), nil
		})
		if err != nil {
			return errors.New("error extract token")

		}
		if !extract.Valid {
			return errors.New("invalid")
		}

		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Successfully logged in", http.StatusOK, responseToken))
	}
}

func (ah *AuthHandler) LoginBuyerHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login entities.Seller
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Error to bind data", http.StatusBadRequest))
		}
		token, errorLogin := ah.authUseCase.LoginBuyer(login.Email, login.Password)

		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed(fmt.Sprintf("%v", errorLogin), http.StatusBadRequest))
		}
		extract, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("S3CR3T"), nil
		})
		if err != nil {
			return errors.New("error extract token")

		}
		if !extract.Valid {
			return errors.New("invalid")
		}

		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("Successfully logged in", http.StatusOK, responseToken))
	}
}
