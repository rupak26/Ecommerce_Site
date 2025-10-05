package user_handler

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"ecommerce/config"
)

type ReqLogin struct {
	Email string      `json:"email"`
	Password string   `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter , r *http.Request) {
	
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&reqLogin) 
    key := config.GetConfig().SecretKey

	if err != nil {
		http.Error(w,"Invalid Request Data" , http.StatusBadRequest)
		return 
	}

	user , err := h.userRepo.Find(reqLogin.Email , reqLogin.Password) 
	if err != nil {
		http.Error(w , "Invalid Credentials" , http.StatusBadRequest)
		return 
	}

	jwt , err := utils.CreateJwt(key , utils.Payload{
		Sub: 1,
		Email: user.Email,
		Password: reqLogin.Password,
	})
	
		
	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Internal Server Error")
		return 
	}
	utils.WriteResponse(w , http.StatusCreated , jwt)
}




