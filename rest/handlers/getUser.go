package handlers

import (
	"net/http"
	"ecommerce/utils"
	"ecommerce/database"
)

func GetUser(w http.ResponseWriter , r *http.Request) {
    utils.WriteResponse(w , http.StatusOK , database.List())
}
