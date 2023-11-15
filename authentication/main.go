package main

import (
	"go-auth-mux/controllers/authcontroller"
	"go-auth-mux/controllers/productcontroller"
	"go-auth-mux/middlewares"
	"go-auth-mux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":5000", r))
}
