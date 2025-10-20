package product_handler

import (
	"ecommerce/utils"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct (w http.ResponseWriter , r *http.Request) {
     id := r.PathValue("id") 
    
	numId , err := strconv.Atoi(id)

	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid user id")
		return
	}
    
	err = h.svc.Delete(numId)

	if err != nil {
		utils.Send_erros(w , "Product not found" , http.StatusNotFound)
		return
	}
	slog.Info("Product Deleted")
	utils.WriteResponse(w , http.StatusOK , "Successfull Deleted")
}