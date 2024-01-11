package main

import (
	"net/http"
	// "strings"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
