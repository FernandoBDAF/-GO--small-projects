package api

import (
	"encoding/json"
	"museum/data"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id != "" {
		for _, exhibition := range data.GetAll() {
			if exhibition.Color == id {
				json.NewEncoder(w).Encode(exhibition)
				return
			}
		}
		http.Error(w, "Exhibition not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data.GetAll())
}