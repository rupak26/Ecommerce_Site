package product_handler

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter , r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit") 
	
	pg, _ := strconv.ParseInt(page, 10, 64)
    lmt, _ := strconv.ParseInt(limit, 10, 64)
	if pg == 0 {
		pg = 1 
	}
	if lmt == 0 {
		lmt = 10
	}
    // http.Error(w , "Internal Server Error" ,http.StatusInternalServerError)

	productListcL := make(chan []domain.Product) 
	countcL    := make(chan int64) 
	errcL       := make(chan error , 2)
    
    go func() {
       productlist , err := h.svc.List(pg ,lmt) 
	   if err != nil {
           errcL <- err  
	   }
       productListcL <- *productlist 
	}()


	go func() {
        totalCount , err := h.svc.Count()
		if err != nil {
		  errcL <- err
		}    
		countcL <- totalCount
	}()
    
	var (
		productList []domain.Product
		totalCount  int64
	)
    
    for i:=0 ; i<2 ; i++ {
		select {
		case products:= <-productListcL:
		     productList = products 
		case count := <-countcL:
			 totalCount = count
		case err := <- errcL:
			 slog.Error(err.Error())
			 http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			 fmt.Println("Error:", err)
			 return
		}
	}

    utils.SendPage(w , productList , pg , lmt , totalCount)
}
