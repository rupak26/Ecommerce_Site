package handlers

import (
	"ecommerce/utils"
	"fmt"
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
	if len(UserList) < numId {
		utils.WriteResponse(w , http.StatusNotFound , fmt.Sprintf("User with id %d not found", numId))
		return 
	}
	utils.WriteResponse(w , http.StatusOK , UserList[numId-1])
}