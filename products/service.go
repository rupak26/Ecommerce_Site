package products


import (
	"ecommerce/domain"
)

type service struct {
	prodRepo ProductRepo
}

func NewService(prodRepo ProductRepo) Service {
    return &service{
		prodRepo: prodRepo,
	}
}

func (svc *service) Create(product domain.Product) (*domain.Product, error) {
     prod , err := svc.prodRepo.Create(product) 
	 if err != nil {
		return nil , err
	 }
	 return prod , nil
}

func (svc *service) Delete(id int) error {
	 err := svc.prodRepo.Delete(id) 
	 if err != nil {
		return err
	 }
	 return nil 
} 

func (svc *service) GetById(pID int) (*domain.Product, error) {
	 prod , err := svc.prodRepo.GetById(pID) 
	 if err != nil {
		return nil , err
	 }
	 return prod , nil
}

func (svc *service) List() (*[]domain.Product, error) {
     prod , err := svc.prodRepo.List()
	 if err != nil {
		return nil , err
	 }
	 return prod , nil
}

func (svc *service) PatchProduct(id int, req domain.UpdateProductReq) (*domain.Product, error) {
	 prod , err := svc.prodRepo.PatchProduct(id , req) 
	 if err != nil {
		return nil , err
	 }
	 return prod , nil
}

func (svc *service) Update(prod domain.Product) (*domain.Product, error) {
	 updatedProd, err := svc.prodRepo.Update(prod)
	 if err != nil {
		return nil, err
	 }
	 return updatedProd, nil
}