package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/golang-gorm-api/db"
	"github.com/julianNot/golang-gorm-api/models"
	"github.com/julianNot/golang-gorm-api/routes"
)

func main() {
	db.DBConnection()
	
	// Create Tables with migrations
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
