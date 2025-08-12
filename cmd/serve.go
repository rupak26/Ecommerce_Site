package cmd


import (
	"fmt"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func Serve() {
	mux := http.NewServeMux() 

	mux.Handle("GET /users" ,middleware.CorsMiddleWare(http.HandlerFunc(handlers.GetUser) ))

	mux.Handle("GET /users/{id}",middleware.CorsMiddleWare(http.HandlerFunc(handlers.GetUserById)))

	mux.Handle("POST /users" , middleware.CorsMiddleWare(http.HandlerFunc(handlers.CreateUser)))

	fmt.Println("The server is running in port 8080") 

	err := http.ListenAndServe(":8080" , mux)

	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}