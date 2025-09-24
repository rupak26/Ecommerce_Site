package rest

import (
	"net/http"
	"ecommerce/rest/middleware"
	"strconv"
	"ecommerce/config"
	"ecommerce/rest/routes"
	"fmt"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager() 

	mux := http.NewServeMux() 

    //replacement 
	routes.InitUserRouter(mux , manager)
	routes.InitProductRouter(mux , manager)
	//mux.Handle("GET /users" ,middleware.Logger(http.HandlerFunc(handlers.GetUser)))
	//mux.Handle("GET /users/{id}",middleware.Logger(http.HandlerFunc(handlers.GetUserById)))
	//mux.Handle("POST /users" ,middleware.Logger(http.HandlerFunc(handlers.CreateUser)))
    
	addr := ":" + strconv.Itoa(int(cnf.HttpPort))
    
	fmt.Println("The server is running in port 8080" , addr) 


	err := http.ListenAndServe(addr , mux)
    
	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}