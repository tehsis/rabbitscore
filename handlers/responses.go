package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type PositionResponse struct {
	Username string `json:"username"`
	Position uint   `json:"position"`
}

func ResponseError(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ErrorResponse{
		message,
	}); err != nil {
		return err
	}

	return nil
}

func ResponsePosition(w http.ResponseWriter, username string, position uint) error {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(PositionResponse{
		username,
		position,
	}); err != nil {
		return err
	}

	return nil
}
