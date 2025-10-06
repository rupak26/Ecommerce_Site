package users

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/user_handler"
)

type Service interface {
    user_handler.Service //embedding
}


type UserRepo interface {
	Create(u domain.User) (*domain.User , error)
	Get() ([]domain.User, error)
	Find(email , password string) (*domain.User , error)
	GetById(usrID int) (*domain.User, error)
	UpdateUser(user domain.User) (*domain.User , error) 
	PatchUser(id int, req domain.UpdateUserReq) (*domain.User, error)
	Delete(id int) (error)
}
