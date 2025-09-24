package routes

import (
	"net/http"
	"ecommerce/rest/handlers/user_handler"
	"ecommerce/rest/middleware"
)

func InitUserRouter(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"GET /users" ,
		manager.With(
			http.HandlerFunc(user_handler.GetUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
    mux.Handle(
		"GET /users/{id}" , 
		manager.With(
			http.HandlerFunc(user_handler.GetUserById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
	mux.Handle(
		"POST /users" ,
		manager.With(
			http.HandlerFunc(user_handler.CreateUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
	)
}