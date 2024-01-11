package main

import (
	"net/http"
	// "sync/atomic"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/chat", authenticate(http.HandlerFunc(handleChat)))

	http.ListenAndServe(":8080", router)
}
