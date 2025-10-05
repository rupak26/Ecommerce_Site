package product_handler

import (
	"net/http"
	"ecommerce/rest/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"POST /products" ,
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
	    ),
	)
	mux.Handle(
		"GET /products" ,
		manager.With(
			http.HandlerFunc(h.GetProducts),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
	    ),
	)
	mux.Handle(
		"GET /products/{id}" ,
		manager.With(
			http.HandlerFunc(h.GetProductById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
	    ),
	)
	mux.Handle(
		"PUT /products/{id}" ,
		manager.With(
			http.HandlerFunc(h.UpdateProductById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
		),
	)
	mux.Handle(
		"PATCH /products/{id}" ,
		manager.With(
			http.HandlerFunc(h.UpdateProductPartialiById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
		),
	)
	mux.Handle(
		"DELETE /products/{id}" ,
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
			h.middleware.Authorization,
		),
	)
}