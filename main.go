package main

import (
	"fmt"
	"net/http"
	"ecommerce/handlers"
)




func corsMiddleWare(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter , r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       w.Header().Set("Acess-Control-Allow-Origin", "*")

	   next.ServeHTTP(w , r)
	}
	handler := http.HandlerFunc(handleCors) 
	return handler
}

func main() {
	mux := http.NewServeMux() 

	mux.Handle("GET /user" ,corsMiddleWare(http.HandlerFunc(handlers.GetUser) ))

	mux.Handle("POST /user" , corsMiddleWare(http.HandlerFunc(handlers.CreateUser)))

	fmt.Println("The server is running in port 8080") 

	err := http.ListenAndServe(":8080" , mux)

	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}