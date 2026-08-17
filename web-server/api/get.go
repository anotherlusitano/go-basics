package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"server.web/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query()["id"]

	if id == nil {
		json.NewEncoder(w).Encode(data.GetAll())
		return
	}

	finalId, err := strconv.Atoi(id[0])

	if err == nil && finalId < len(data.GetAll()) {
		json.NewEncoder(w).Encode(data.GetAll()[finalId])
	} else {
		http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
	}
}
