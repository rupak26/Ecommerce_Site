package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product_handler"
	"ecommerce/rest/handlers/user_handler"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	// Storing as Dependecies 
	cnf            *config.Config
	productHandler *product_handler.Handler
	userHandler    *user_handler.Handler
}

func NewServer(
	cnf            *config.Config,
	productHandler *product_handler.Handler,
	userHandler    *user_handler.Handler,
) *Server {
	return &Server{
		 cnf            : cnf,
         productHandler : productHandler ,
		 userHandler    : userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager() 

	mux := http.NewServeMux() 


	server.productHandler.RegisterRouters(mux , manager)
	server.userHandler.RegisterRouters(mux , manager)
    //replacement 
	
	//mux.Handle("GET /users" ,middleware.Logger(http.HandlerFunc(handlers.GetUser)))
	//mux.Handle("GET /users/{id}",middleware.Logger(http.HandlerFunc(handlers.GetUserById)))
	//mux.Handle("POST /users" ,middleware.Logger(http.HandlerFunc(handlers.CreateUser)))
    
	addr := ":" + strconv.Itoa(int(server.cnf.HttpPort))
    
	fmt.Println("The server is running in port 8080" , addr) 


	err := http.ListenAndServe(addr , mux)
    
	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}