package main

import (
	"encoding/json"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct {
		Version string
		Status  string `json:"status"`
	}{
		"1.4",
		"OK",
	}); err != nil {
		panic(err)
	}
}
