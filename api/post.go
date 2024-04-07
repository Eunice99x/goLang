package api

import (
	"encoding/json"
	"net/http"

	"younes.dev/go/data"
)

func Post(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		var item data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&item)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(item)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else{
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}