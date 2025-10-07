package product_handler

import (
	"net/http"
	"ecommerce/utils"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter , r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit") 
	
	pg, _ := strconv.ParseInt(page, 10, 64)
    lmt, _ := strconv.ParseInt(limit, 10, 64)

	productlist , err := h.svc.List(pg ,lmt) 
	
	if err != nil {
       http.Error(w , "Internal Server Error" ,http.StatusInternalServerError)
	}
    utils.WriteResponse(w , http.StatusOK , productlist)
}
