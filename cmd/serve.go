package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager() 

	mux := http.NewServeMux() 

    //replacement 
	InitRouter(mux , manager)
	//mux.Handle("GET /users" ,middleware.Logger(http.HandlerFunc(handlers.GetUser)))

	//mux.Handle("GET /users/{id}",middleware.Logger(http.HandlerFunc(handlers.GetUserById)))

	//mux.Handle("POST /users" ,middleware.Logger(http.HandlerFunc(handlers.CreateUser)))
    
	addr := ":" + string(cnf.HttpPort)

	fmt.Println("The server is running in port 8080" , addr) 

	globalrouter := middleware.CorsWithPrefight(mux)

	err := http.ListenAndServe(addr , globalrouter)
    
	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}