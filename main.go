package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/encurtador-url-go/config"
)

type Relation struct {
	Url, Code string
}

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
	r.HandleFunc("/desencurta", Desencurta).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func Encurta(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	code := RandStringRunes(6)
	
	db := config.GetConexao()
	defer db.Close()

	//I should test if this sentence really get the existing code without any treatment in the result
	existingCode, err := db.Query(fmt.Sprintf("SELECT * FROM relation WHERE code = '%s'", code))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	//Must test if this first case calls the method itself
	if existingCode != nil {
		Encurta(w, r)
	}else{
		relation := Relation{
			Url:  url,
			Code: code,
		}
	
		_, err = db.Exec(fmt.Sprintf("INSERT INTO relation (url, code) VALUES ('%s','%s')", relation.Url, relation.Code))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
	
	
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(relation.Code)
	}
}

func Desencurta(w http.ResponseWriter, r *http.Request) {
	
}