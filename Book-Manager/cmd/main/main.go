package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeremiah-quaynor/golang-projects/Book-Manager/pkg/routes"
	"github.com/jeremiah-quaynor/golang-projects/pkg/rountes"
	_ "github.com/jinzhu/gorm/diablects/myqsl"
)

func main ( ) {
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}