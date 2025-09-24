package product_handler

import (
	"net/http"
	"ecommerce/utils"
	"ecommerce/database/product_database"
)

func GetProducts(w http.ResponseWriter , r *http.Request) {
    utils.WriteResponse(w , http.StatusOK , productdatabase.GetProductList())
}
