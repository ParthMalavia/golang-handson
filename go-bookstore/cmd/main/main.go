package main

import (
	"log"
	"net/http"

	"example/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r) // Comment-out
	log.Fatal(http.ListenAndServe(":9010", r))
}
