package user_handler

import "ecommerce/repo"

type Handler struct {
     userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}