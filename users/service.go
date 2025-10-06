package users

import (
	"ecommerce/domain"
)

type service struct {
	usrRepo UserRepo 
}

func NewService(usrRepo UserRepo) Service {
    return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr , err := svc.usrRepo.Create(user)
	if err != nil {
		return nil , err
	}
	if usr == nil {
		return nil , nil
	}
	return usr , nil 
}

func (svc *service) Delete(id int) error {
	err := svc.usrRepo.Delete(id) 
	if err != nil {
		return err
	}
	return nil
}
func (svc *service) Find(email string , password string) (*domain.User , error) {
	usr , err := svc.usrRepo.Find(email , password) 
	if err != nil {
		return nil , err
	}
	return usr , nil
}
func (svc *service) Get() ([]domain.User , error) {
	usr , err := svc.usrRepo.Get() 
	if err != nil {
		return nil , err
	}
	return usr , nil
}
func (svc *service) GetById(usrId int) (*domain.User , error) {
	usr , err := svc.usrRepo.GetById(usrId) 
	if err != nil {
		return nil , err
	}
	return usr , nil
}
func (svc *service) PatchUser(id int, req domain.UpdateUserReq) (*domain.User, error) {
	usr , err := svc.usrRepo.PatchUser(id , req) 
	if err != nil {
		return nil , err
	}
	return usr , nil
}
func (svc *service) UpdateUser(user domain.User) (*domain.User, error) {
	usr , err := svc.usrRepo.UpdateUser(user)
	if err != nil {
		return nil , err
	}
	return usr , nil
}
