package handlers

import (
	"net/http"
	"ecommerce/utils"
)

func GetUser(w http.ResponseWriter , r *http.Request) {
    utils.WriteResponse(w , http.StatusOK , UserList)
}
