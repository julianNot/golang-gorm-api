package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/routes"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
