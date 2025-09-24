package user_handler

import (
	"ecommerce/database/user_database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)



func Login(w http.ResponseWriter , r *http.Request) {
	
	var reqLogin userdatabase.ReqLogin

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&reqLogin) 
    key := "a4bc166839d6ff3c83eb9d1cbb0ddda2a65f74eeea9a07413357200dfad1d808d708e7d5"

	if err != nil {
		http.Error(w,"Invalid Request Data" , http.StatusBadRequest)
		return 
	}
    
	jwt , err := utils.CreateJwt(key , utils.Payload{
		Sub: newUser.Id,
		Name: newUser.Name,
		Email: newUser.Email,
	})

	user := userdatabase.Find(reqLogin.Email , reqLogin.Password) 
	if user == nil {
		http.Error(w , "Invalid Credentials" , http.StatusBadRequest)
		return 
	}
	
	utils.WriteResponse(w , http.StatusAccepted , user)
}




// 	userdatabase.StoreUser(newUser)
// 	if err != nil {
// 		utils.WriteResponse(w , http.StatusBadRequest , "Invalid")
// 	}
// 	utils.WriteResponse(w , http.StatusCreated , jwt)