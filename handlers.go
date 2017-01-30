package main

import (
	"encoding/json"
	"net/http"

	"github.com/tehsis/rabbitscore/services/redis"
)

func Status(w http.ResponseWriter, r *http.Request) {
	redisClient := redis.GetClient()
	var status string

	pong, _ := redisClient.Ping().Result()

	if pong == "PONG" {
		w.WriteHeader(http.StatusOK)
		status = "OK"
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		status = "FAILED"
	}

	if err := json.NewEncoder(w).Encode(struct {
		Version string `json:"version"`
		Status  string `json:"status"`
	}{
		"1.4",
		status,
	}); err != nil {
		panic(err)
	}
}
