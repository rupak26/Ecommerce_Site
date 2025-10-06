package product_handler


import (
	"net/http"
	"strconv"
	"ecommerce/utils"
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
	utils.WriteResponse(w , http.StatusOK , "Successfull Deleted")
}