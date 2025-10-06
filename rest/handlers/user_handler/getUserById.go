package user_handler

import (
	"ecommerce/utils"
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
	
	user , err := h.svc.GetById(numId)
	if err != nil {
		utils.Send_erros(w , "User not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , user)
}