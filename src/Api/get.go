package api

import (
	"encoding/json"
	"golang_api/src/data"
	"io"
	"net/http"
)

func GETHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	rez := data.GetData(string(body))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rez)
}
