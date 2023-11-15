package main

import (
	"go-api-mux/controllers/employecontroller"
	"go-api-mux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/employes", employecontroller.Index).Methods("GET")
	r.HandleFunc("/employe/{id}", employecontroller.Show).Methods("GET")
	r.HandleFunc("/employe", employecontroller.Create).Methods("POST")
	r.HandleFunc("/employe/{id}", employecontroller.Update).Methods("PUT")
	r.HandleFunc("/employe", employecontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
