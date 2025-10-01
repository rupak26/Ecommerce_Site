package routes

import (
	"net/http"
	"ecommerce/rest/handlers/product_handler"
	"ecommerce/rest/middleware"
)

func InitProductRouter(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"POST /products" ,
		manager.With(
			http.HandlerFunc(product_handler.CreateProduct),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			middleware.Authorization,
	    ),
	)
	mux.Handle(
		"GET /products" ,
		manager.With(
			http.HandlerFunc(product_handler.GetProducts),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			middleware.Authorization,
	    ),
	)
	mux.Handle(
		"GET /products/{id}" ,
		manager.With(
			http.HandlerFunc(product_handler.GetProductById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			middleware.Authorization,
	    ),
	)
}