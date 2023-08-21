package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lostinopensrc/go-crud-api/schema"
)

func Get_Movies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schema.Movies)
}

func Delete_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // parameters from request
	for index, item := range schema.Movies {
		if item.Id == params["id"] {
			schema.Movies = append(schema.Movies[:index], schema.Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(schema.Movies)
}

func Get_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // parameters from request
	for _, item := range schema.Movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func Create_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie schema.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	schema.Movies = append(schema.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}
