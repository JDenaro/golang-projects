package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server started on: http://localhost:9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
