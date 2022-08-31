package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grzelkowska/11projects/03go_bookstore_mysql/pkg/routes"
	// "github.com/grzelkowska/go03_bookstore_mysql/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}