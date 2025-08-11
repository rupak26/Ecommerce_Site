package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string 
	Age  int 
	Occupation string
}

var UserList[] User

func writeResponse(w http.ResponseWriter , status int , data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func getUser(w http.ResponseWriter , r *http.Request) {
	
    writeResponse(w , http.StatusOK , UserList)
}


func createUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser User

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}

	writeResponse(w , http.StatusCreated , newUser)
	UserList = append(UserList, newUser)
}

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
    
	handler := http.HandlerFunc(getUser) 
    updatedHandler := corsMiddleWare(handler)

	handler2 := http.HandlerFunc(createUser) 
    updatedHandler2 := corsMiddleWare(handler2)

	mux.Handle("GET /user" ,updatedHandler)

	mux.Handle("POST /user" , updatedHandler2)

	fmt.Println("The server is running in port 8080") 

	err := http.ListenAndServe(":8080" , mux)

	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}