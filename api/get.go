package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"younes.dev/go/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	queryString := r.URL.Query()
	id := queryString["id"]

	if id != nil {
		index, err := strconv.Atoi(id[0])

		if err == nil && index < len(data.GetAll()) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data.GetAll()[index])
		}else{
			http.Error(w, "Invalid exhibition", http.StatusBadRequest)
		}
	}else {
		json.NewEncoder(w).Encode(data.GetAll())
	}

}