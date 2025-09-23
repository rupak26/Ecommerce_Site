package rest

import (
	"net/http"
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
)

func InitRouter(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"GET /users" ,
		manager.With(
			http.HandlerFunc(handlers.GetUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
    mux.Handle(
		"GET /users/{id}" , 
		manager.With(
			http.HandlerFunc(handlers.GetUserById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
	mux.Handle(
		"POST /users" ,
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
	)
}