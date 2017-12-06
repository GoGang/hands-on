package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func (t *Talk) AddProperty(name string, value interface{}) *Talk {
	t.Properties[name] = value
	return t
}

func NewTalk(name, description string) *Talk {
	return &Talk{
		Name:        name,
		Description: description,
		Properties:  make(map[string]interface{}),
	}
}

func main() {
	talks := []*Talk{ // on remplit un peu notre "repository"
		&Talk{"Jacques-Antoine", "XMS", map[string]interface{}{}},
		&Talk{"Thérèse", "SOS Amitié", map[string]interface{}{}},
	}

	// on crée quelques contrôleurs pour un API REST de base
	http.HandleFunc("/api/talks/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		e.Encode(talks[1])
	})
	http.HandleFunc("/api/talks/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		e.Encode(talks)
	})

	// on démarre le serveur en écrivant des logs
	log.Fatal(http.ListenAndServe(":8666", nil))
}
