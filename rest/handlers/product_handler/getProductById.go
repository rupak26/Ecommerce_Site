package product_handler

import (
	"ecommerce/utils"
	"ecommerce/database/product_database"
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
	product := productdatabase.GetProduct(prodId) 
	if product == nil {
		utils.Send_erros(w , "Product not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , productdatabase.GetProduct(prodId))
}