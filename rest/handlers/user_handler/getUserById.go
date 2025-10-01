package user_handler

import (
	"ecommerce/utils"
	"ecommerce/database/user_database"
	//"fmt"
	"net/http"
	"strconv"
)


func (h *Handler) GetUserById(w http.ResponseWriter , r *http.Request) {
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
	user := userdatabase.GetUser(numId) 
	if user == nil {
		utils.Send_erros(w , "User not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , userdatabase.GetUser(numId))
}