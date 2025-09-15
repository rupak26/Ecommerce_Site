package cmd 

import (
	"net/http"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func InitRouter(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"GET /users" ,
		manager.With(
			http.HandlerFunc(handlers.GetUser),
			middleware.Logger,
	    ),
    )
    mux.Handle(
		"GET /users/{id}" , 
		manager.With(
			http.HandlerFunc(handlers.GetUserById),
			middleware.Logger,
	    ),
    )
	mux.Handle(
		"POST /users" ,
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
			middleware.Logger,
	    ),
	)
}