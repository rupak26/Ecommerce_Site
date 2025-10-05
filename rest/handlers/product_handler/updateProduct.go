package product_handler

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateProductReq struct {
	ProductName *string `json:"productname"`
	Url  *string        `json:"url"`
	Quantity *int       `json:"quantity"`
}

func (h *Handler) UpdateProductById(w http.ResponseWriter , r *http.Request) {
	id := r.PathValue("id") 
    
	var req UpdateProductReq

	numId , err := strconv.Atoi(id)

	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid user id")
		return
	}
    
	decoder := json.NewDecoder(r.Body) 
	err = decoder.Decode(&req) 
    
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
    
	defer r.Body.Close()
    
	if req.ProductName == nil && req.Quantity == nil && req.Url == nil {
		utils.WriteResponse(w, http.StatusBadRequest, "No fields provided for update")
		return
    }

    prod := repo.Product {
		Id:         numId,
		ProductName : "" ,
		Url : "" ,
		Quantity : 0,
	}
	

	if req.ProductName != nil {
		prod.ProductName = *req.ProductName
	}
	if req.Url != nil {
		prod.Url = *req.Url
	}
	if req.Quantity != nil {
	   prod.Quantity = *req.Quantity
	}
	
	updatedUser , err := h.productRepo.Update(prod)
	if err != nil {
		utils.Send_erros(w , "User not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , updatedUser)
}