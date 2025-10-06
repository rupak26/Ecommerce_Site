package user_handler

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserReq struct {
	Name string         `json:"name"`
	Age  int            `json:"age"`
	Email string        `json:"email"`
	Password string     `json:"password"`
	Occupation string   `json:"occupation"`
}

func (h *Handler) CreateUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser CreateUserReq

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 
    
	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}
    
	createUser , err := h.svc.Create(domain.User{
		Name: newUser.Name,
		Age: newUser.Age,
		Email: newUser.Email,
		Password: newUser.Password,
		Occupation: newUser.Occupation,
	})
	
	if err != nil {
		http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
		return
	}
	
	
	utils.WriteResponse(w , http.StatusCreated , createUser)
}
