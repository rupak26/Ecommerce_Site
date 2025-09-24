package product_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/utils"
	"ecommerce/database/product_database"
)



func CreateProduct(w http.ResponseWriter , r *http.Request) {
	
	var newProduct productdatabase.Product

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newProduct) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}

	newProduct.Id = len(productdatabase.ProductList) + 1

	utils.WriteResponse(w , http.StatusCreated , newProduct)

	productdatabase.StoreProduct(newProduct)
}
