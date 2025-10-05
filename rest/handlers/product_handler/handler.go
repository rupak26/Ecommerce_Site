package product_handler

import (
	 "ecommerce/rest/middleware"
	 "ecommerce/repo"
)

type Handler struct {
	middleware *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
    
	) *Handler {
		
	return &Handler{
		middleware: middlewares,
		productRepo: productRepo,
	}
}