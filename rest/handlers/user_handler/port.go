package user_handler

import (
	"ecommerce/domain"
)

type Service interface {
	Create(u domain.User) (*domain.User , error)
	Get() ([]domain.User, error)
	Find(email , password string) (*domain.User , error)
	GetById(usrID int) (*domain.User, error)
	UpdateUser(user domain.User) (*domain.User , error) 
	PatchUser(id int, req domain.UpdateUserReq) (*domain.User, error)
	Delete(id int) (error)
}


