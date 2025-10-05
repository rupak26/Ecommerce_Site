package product_handler

import (

	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	//"hash"
	"net/http"
	//"github.com/go-playground/locales/tk"
)

type RequestCreateProduct struct {
	ProductName string `json:"productname"`
	Url  string        `json:"url"`
	Quantity int       `json:"quantity"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter , r *http.Request) {

	var newProduct RequestCreateProduct

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newProduct) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}

	createProduct , err := h.productRepo.Create(repo.Product{
		ProductName : newProduct.ProductName ,
		Url: newProduct.Url,
		Quantity: newProduct.Quantity,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.WriteResponse(w , http.StatusCreated , createProduct)
}
