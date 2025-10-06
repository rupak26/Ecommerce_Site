package user_handler

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateUserReq struct {
	Name *string         `json:"name"`
	Age  *int            `json:"age"`
	Email *string        `json:"email"`
	Password *string     `json:"password"`
	Occupation *string   `json:"occupation"`
}

func (h *Handler) UpdateUserById(w http.ResponseWriter , r *http.Request) {
	id := r.PathValue("id") 
    
	var req UpdateUserReq

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
    
	if req.Name == nil && req.Age == nil && req.Email == nil && req.Occupation == nil && req.Password == nil {
		utils.WriteResponse(w, http.StatusBadRequest, "No fields provided for update")
		return
    }


	user := domain.User{
		Id:         numId,
		Name:       "",
		Age:        0,
		Email:      "",
		Password:   "",
		Occupation: "",
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Age != nil {
		user.Age = *req.Age
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
	if req.Occupation != nil {
		user.Occupation = *req.Occupation
	}

	updatedUser , err := h.svc.UpdateUser(user)
	if err != nil {
		utils.Send_erros(w , "User not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , updatedUser)
}