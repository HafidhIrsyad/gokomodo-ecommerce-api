package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gokomodo-api/configs"
	"gokomodo-api/delivery/handler"
	_middlerware "gokomodo-api/delivery/middlerwares"
	"gokomodo-api/delivery/router"
	"gokomodo-api/driver"
	_authRepo "gokomodo-api/repository/auth"
	_orderRepo "gokomodo-api/repository/order"
	_productRepo "gokomodo-api/repository/product"
	_userRepo "gokomodo-api/repository/user"
	_authService "gokomodo-api/service/auth"
	_orderService "gokomodo-api/service/order"
	_productService "gokomodo-api/service/product"
	_userService "gokomodo-api/service/user"
	"log"
	"net/http"
)

func main() {
	configsDb := configs.GetConfig()
	db := driver.InitDB(configsDb)

	userRepo := _userRepo.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authRepo := _authRepo.NewAuthRepository(db)
	authService := _authService.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	productRepo := _productRepo.NewFacilityRepository(db)
	productService := _productService.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	orderRepo := _orderRepo.NewOrderRepository(db)
	orderService := _orderService.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(_middlerware.CustomLogger())

	router.RegisterUserPath(e, userHandler)
	router.RegisterAuthPath(e, authHandler)
	router.ProductPath(e, productHandler)
	router.OrderPath(e, orderHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", configsDb.Port)))

}
