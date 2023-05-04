package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/encurta", Encurta).Methods("POST")
	r.HandleFunc("/desencurta", Desencurta).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func Encurta(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}

func Desencurta(w http.ResponseWriter, r *http.Request) {

}
