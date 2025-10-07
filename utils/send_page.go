package utils

import (
	"net/http"
)

type PaginatedData struct {
    DataList    any `json:"datalist"`
	Pagination  Pagination
}

type Pagination struct {
	Page       int64             `json:"page"`
	Limit      int64             `json:"limit"`
	TotalItems int64             `json:"totalitems"`
    TotalPages int64             `json:"totalpages"`
} 


func SendPage(w http.ResponseWriter , data any  , page , limit , totalcount int64)  {
     paginatedData := PaginatedData{
		DataList: data,
		Pagination: Pagination{
			Page: page,
			Limit: limit,
			TotalItems: totalcount,
			TotalPages: totalcount / limit,
		},
	 }
	 WriteResponse(w , http.StatusOK , paginatedData)
}