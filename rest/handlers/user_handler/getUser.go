package user_handler

import (
	"net/http"
	"ecommerce/utils"
	"ecommerce/database/user_database"
)

func GetUser(w http.ResponseWriter , r *http.Request) {
    utils.WriteResponse(w , http.StatusOK , userdatabase.GetUserList())
}
