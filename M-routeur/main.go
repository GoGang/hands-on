package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // go get pour récupérer la lib dans votre GOPATH
)

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func toJson(v interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v) // Marshalling JSON
}

func main() {
	talks := []*Talk{ // on remplit un peu notre "repository"
		&Talk{"Jacques-Antoine", "XMS", map[string]interface{}{}},
		&Talk{"Thérèse", "SOS Amitié", map[string]interface{}{}},
	}

	r := mux.NewRouter()

	// on écrit quelques contrôleurs
	r.HandleFunc("/api/talks/{id}", func(w http.ResponseWriter, r *http.Request) { // extrayez les fonctions pour rendre le code plus lisible
		params := mux.Vars(r)
		pos, _ := strconv.Atoi(params["id"])
		toJson(talks[pos], w)
	}).Methods("GET")

	r.HandleFunc("/api/talks", func(w http.ResponseWriter, r *http.Request) {
		toJson(talks, w)
	}).Methods("GET")

	r.HandleFunc("/talks", func(w http.ResponseWriter, r *http.Request) {
		for _, t := range talks {
			w.Write([]byte(fmt.Sprintf(`%s : %s`, t.Name, t.Description)))
		}
	}).Methods("GET")

	// on démarre le serveur
	err := http.ListenAndServe(":8666", r)
	if err != nil {
		log.Println(err)
	}
}
