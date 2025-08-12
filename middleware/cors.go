package middleware

import (
	"net/http"
)


func CorsMiddleWare(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter , r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       w.Header().Set("Acess-Control-Allow-Origin", "*")

	   next.ServeHTTP(w , r)
	}
	handler := http.HandlerFunc(handleCors) 
	return handler
}