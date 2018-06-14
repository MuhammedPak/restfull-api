package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Member struct {
	Name      string
	Surname   string
	Age       int
	Birthdate string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/MemberInfo", Member_Info)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func Member_Info(w http.ResponseWriter, r *http.Request) {

	var err error
	result := Member{"Muhammed", "Pak", 24, "1994"}

	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
