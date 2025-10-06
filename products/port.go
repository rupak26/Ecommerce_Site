package products


import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/product_handler"
)

type Service interface {
	product_handler.Service //embeding
}

type ProductRepo interface{
	Create(p domain.Product) (*domain.Product , error) 
	GetById(pID int) (*domain.Product , error)
	List() (*[]domain.Product , error)
	Update(prod domain.Product) (*domain.Product , error)
	PatchProduct(id int, req domain.UpdateProductReq) (*domain.Product, error)
	Delete(id int) error
} 
