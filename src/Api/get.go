package api

import (
	"encoding/json"
	"golang_api/src/data"
	"net/http"
)

func GETHandler(w http.ResponseWriter, r *http.Request) {
	body := r.URL.Query().Get("words")
	rez := data.GetData(string(body))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rez)
}
