package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/encurtador-url-go/config"
	"github.com/ropehapi/encurtador-url-go/model"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/encurta", Encurta).Methods("POST")
	r.HandleFunc("/desencurta/{code}", Desencurta).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// TODO: I must validate if the code already exists in a relation
func Encurta(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	code := RandStringRunes(6)

	relation := model.Relation{
		Url:  url,
		Code: code,
	}

	err := relation.Store()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(relation.Code)
}

func Desencurta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	db := config.GetConexao()
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT url FROM relation WHERE code = '%s'", code))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	var url string
	res.Next()
	res.Scan(&url)
	if err := res.Scan(&url); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Code not found")
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}
