package product_handler

import (
	"net/http"
	"ecommerce/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter , r *http.Request) {
	productlist , err := h.productRepo.List() 
	
	if err != nil {
       http.Error(w , "Internal Server Error" ,http.StatusInternalServerError)
	}
    utils.WriteResponse(w , http.StatusOK , productlist)
}
