package main

import (
	"log"
	"net/http"

	"github.com/grzelkowska/11projects/03go_bookstore_mysql/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
 