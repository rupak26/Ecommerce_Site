package user_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/utils"
	"ecommerce/database/user_database"
)



func CreateUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser userdatabase.User

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 
    
	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}
    
	
	newUser.Id = len(userdatabase.UserList) + 1

	userdatabase.StoreUser(newUser)
	
	utils.WriteResponse(w , http.StatusCreated , newUser)
}
