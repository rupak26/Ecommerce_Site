package user_handler

import (
	"net/http"
	"ecommerce/rest/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"GET /users" ,
		manager.With(
			http.HandlerFunc(h.GetUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
    mux.Handle(
		"GET /users/{id}" , 
		manager.With(
			http.HandlerFunc(h.GetUserById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
    )
	mux.Handle(
		"POST /users" ,
		manager.With(
			http.HandlerFunc(h.CreateUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
	)
	mux.Handle(
		"POST /users/login" ,
		manager.With(
			http.HandlerFunc(h.Login),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
	    ),
	)
	mux.Handle(
		"PUT /users/{id}" ,
		manager.With(
			http.HandlerFunc(h.UpdateUserById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
		),
	)
	mux.Handle(
		"PATCH /users/{id}" ,
		manager.With(
			http.HandlerFunc(h.UpdateUserPatialById),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
		),
	)
	mux.Handle(
		"DELETE /users/{id}" ,
		manager.With(
			http.HandlerFunc(h.DeleteUser),
			middleware.Logger,
			middleware.Cors,
			middleware.Prefight,
		),
	)
}