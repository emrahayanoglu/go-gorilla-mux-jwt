package main

import (
	"fmt"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		if token, ok := r.Header["Authorization"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Authorization Header Missing")
		} else {
			if len(token) < 1 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Authorization Token Not Available")
			} else {
				_, err := ValidateToken(token[0])
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Authorization Token Not Valid")
				} else {
					fmt.Println("Test Middleware is executed!")
					// Call the next handler, which can be another middleware in the chain, or the final handler.
					next.ServeHTTP(w, r)
				}
			}
		}
	})
}
