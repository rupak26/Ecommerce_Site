package middleware

import (
	"net/http"
)


func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		// handle cors
       w.Header().Set("Content-Type", "application/json")
       w.Header().Set("Acess-Control-Allow-Origin", "*")
       
	   next.ServeHTTP(w , r)
	})
}