package handlers

import (
	"ecommerce/utils"
	"ecommerce/database"
	//"fmt"
	"net/http"
	"strconv"
)


func GetUserById(w http.ResponseWriter , r *http.Request) {
    id := r.PathValue("id") 
	 
	numId , err := strconv.Atoi(id)

	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid user id")
		return
	}
	// if len(database.UserList) < numId {
	// 	utils.WriteResponse(w , http.StatusNotFound , fmt.Sprintf("User with id %d not found", numId))
	// 	return 
	// }
	user := database.Get(numId) 
	if user == nil {
		utils.Send_erros(w , "User not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , database.Get(numId))
}