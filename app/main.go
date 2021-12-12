package main

import (
	"fmt"
	"net/http"

	api "github.com/busraarsln/rest-api-with-go/apis"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Listen check-1")

	r := mux.NewRouter()

	r.HandleFunc("/products", api.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", api.GetProducts).Methods("GET")
	r.HandleFunc("/products", api.GetProducts).Methods("POST")
	r.HandleFunc("/products/{id}", api.GetProducts).Methods("PUT")
	r.HandleFunc("/products/{id}", api.GetProducts).Methods("DELETE")
	r.HandleFunc("/categories", api.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", api.GetProducts).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	fmt.Println("Listen check-2")
}
