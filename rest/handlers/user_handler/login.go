package user_handler

import (
	"ecommerce/database/user_database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"ecommerce/config"
)



func Login(w http.ResponseWriter , r *http.Request) {
	
	var reqLogin userdatabase.ReqLogin

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&reqLogin) 
    key := config.GetConfig().SecretKey

	if err != nil {
		http.Error(w,"Invalid Request Data" , http.StatusBadRequest)
		return 
	}

	user := userdatabase.Find(reqLogin.Email , reqLogin.Password) 
	if user == nil {
		http.Error(w , "Invalid Credentials" , http.StatusBadRequest)
		return 
	}

	jwt , err := utils.CreateJwt(key , utils.Payload{
		Sub: 1,
		Email: reqLogin.Email,
		Password: reqLogin.Password,
	})
	
		
	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Internal Server Error")
		return 
	}
	utils.WriteResponse(w , http.StatusCreated , jwt)
}




