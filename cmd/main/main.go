package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishmamabid99/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
