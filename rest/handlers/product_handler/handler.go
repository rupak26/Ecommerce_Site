package product_handler

import "ecommerce/rest/middleware"

type Handler struct {
	middleware *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middleware: middlewares,
	}
}