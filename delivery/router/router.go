package router

import (
	"github.com/labstack/echo/v4"
	_handler "gokomodo-api/delivery/handler"
	"gokomodo-api/delivery/middlerwares"
)

func RegisterAuthPath(e *echo.Echo, ah *_handler.AuthHandler) {
	e.POST("/seller/login", ah.LoginSellerHandler())
	e.POST("/buyer/login", ah.LoginBuyerHandler())
}

func RegisterUserPath(e *echo.Echo, uh _handler.UserHandler) {
	e.POST("/seller/register", uh.CreateSellerHandler())
	e.POST("/buyer/register", uh.CreateBuyerHandler())
}

func ProductPath(e *echo.Echo, ph _handler.ProductHandler) {
	e.GET("/product/list", ph.GetAllProductHandler())
	e.POST("/product", ph.CreateProductHandler(), middlerwares.JWTMiddleware())
}

func OrderPath(e *echo.Echo, oh _handler.OrderHandler) {
	e.GET("/seller/order-list", oh.GetOrderSellerHandler())
	e.GET("/buyer/product-list", oh.GetOrderBuyerHandler())
	e.POST("/order/product", oh.CreateBookingHandler(), middlerwares.JWTMiddleware())
}
