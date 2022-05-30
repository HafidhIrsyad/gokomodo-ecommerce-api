package handler

import (
	"github.com/labstack/echo/v4"
	"gokomodo-api/delivery/middlerwares"
	"gokomodo-api/delivery/response"
	_entities "gokomodo-api/entities"
	"gokomodo-api/service/product"
	"net/http"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func NewProductHandler(c product.ProductServiceInterface) ProductHandler {
	return ProductHandler{
		productService: c,
	}
}

func (ph *ProductHandler) GetAllProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := ph.productService.GetAllProduct()

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("failed get all product", http.StatusBadRequest))
		}

		responseProducts := []map[string]interface{}{}

		for i := 0; i < len(products); i++ {
			response := map[string]interface{}{
				"ID":           products[i].ID,
				"product_name": products[i].ProductName,
				"description":  products[i].Description,
				"price":        products[i].Price,
			}
			responseProducts = append(responseProducts, response)
		}

		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all product", http.StatusOK, responseProducts))
	}
}

func (ph *ProductHandler) CreateProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param _entities.Product

		idToken, errToken := middlerwares.ExtractToken(c)

		if errToken != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("Unauthorized", http.StatusBadRequest))
		}

		param.Seller = idToken

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, err = ph.productService.CreateProduct(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ResponseFailed("created product failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("created product successfully", http.StatusOK))
	}

}
