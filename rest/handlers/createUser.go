package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/utils"
	"ecommerce/database"
)



func CreateUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser database.User

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}

	newUser.Id = len(database.UserList) + 1

	utils.WriteResponse(w , http.StatusCreated , newUser)

	database.Store(newUser)
}
