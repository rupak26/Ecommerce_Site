package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/types"
	"ecommerce/utils"
)

var UserList[] types.User

func CreateUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser types.User

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}

	utils.WriteResponse(w , http.StatusCreated , newUser)
	UserList = append(UserList, newUser)
}
