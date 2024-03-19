package api

import (
	"encoding/json"
	"museum/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
		return
	}
	// id := r.URL.Query().Get("id")
	// if id != "" {
	// 	for _, exhibition := range data.GetAll() {
	// 		if exhibition.Color == id {
	// 			json.NewEncoder(w).Encode(exhibition)
	// 			return
	// 		}
	// 	}
	// 	http.Error(w, "Exhibition not found", http.StatusNotFound)
	// 	return
	// }
	json.NewEncoder(w).Encode(data.GetAll())
}