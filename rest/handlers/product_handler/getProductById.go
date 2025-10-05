package product_handler

import (
	"ecommerce/utils"
	//"fmt"
	"net/http"
	"strconv"
)


func (h * Handler) GetProductById(w http.ResponseWriter , r *http.Request) {
    id := r.PathValue("id") 
	 
	prodId , err := strconv.Atoi(id)

	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid prod id")
		return
	}
	// if len(database.UserList) < numId {
	// 	utils.WriteResponse(w , http.StatusNotFound , fmt.Sprintf("User with id %d not found", numId))
	// 	return 
	// }
	product , err := h.productRepo.Get(prodId)
	if err != nil {
		utils.Send_erros(w , "Product not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , product)
}