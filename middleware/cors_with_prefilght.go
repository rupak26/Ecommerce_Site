package middleware

import (
	"net/http"
)


func CorsWithPrefight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		// handle cors
       w.Header().Set("Content-Type", "application/json")
       w.Header().Set("Acess-Control-Allow-Origin", "*")
       
	   // handles prefight request
	   if r.Method == "OPTIONS" {
		 w.WriteHeader(200) 
		 return 
	   }

	   next.ServeHTTP(w , r)
	})
}