package product_handler

import (
	 "ecommerce/rest/middleware"
)

type Handler struct {
	middleware *middleware.Middlewares
	svc Service
}

func NewHandler(
	middlewares *middleware.Middlewares,
	svc Service,
    
	) *Handler {
		
	return &Handler{
		middleware: middlewares,
		svc: svc,
	}
}