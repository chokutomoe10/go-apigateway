package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Employe struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func employeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author = r.Header.Get("Authorization")

		if author != "employe" {
			w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

func getAllEmployes(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3000/employes")

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var employe = &[]Employe{}
	json.Unmarshal(data, &employe)
	json.NewEncoder(w).Encode(employe)
}

func getOneEmploye(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	resp, _ := http.Get("http://localhost:3000/employe/" + id)

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var employe Employe
	json.Unmarshal(data, &employe)
	json.NewEncoder(w).Encode(employe)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/getallemployes", employeMiddleware(getAllEmployes))
	mux.HandleFunc("/getoneemploye/{id}", getOneEmploye)

	fmt.Println("Server running")
	http.ListenAndServe(":7000", mux)
}
