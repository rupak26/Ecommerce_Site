package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product_handler"
	"ecommerce/rest/handlers/user_handler"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
    
	middleware := middleware.NewMiddleware(cnf)
    productHandler := product_handler.NewHandler(middleware)
	userHandler := user_handler.NewHandler()

    server := rest.NewServer(cnf , productHandler , userHandler)
	server.Start()
}