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

	if err != nil {
		http.Error(w,"Invalid Request Data" , http.StatusBadRequest)
		return 
	}

	user := userdatabase.Find(reqLogin.Email , reqLogin.Password) 
	if user == nil {
		http.Error(w , "Invalid Credentials" , http.StatusBadRequest)
		return 
	}
	
	utils.WriteResponse(w , http.StatusAccepted , user)
}

