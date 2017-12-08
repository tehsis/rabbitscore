package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tehsis/rabbitscore/services/logger"
)

func main() {
	logger.Log().Info("starting server")
	router := NewRouter()
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
