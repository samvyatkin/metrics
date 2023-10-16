package main

import (
	"metrics/internal/database"
	"metrics/internal/handlers/auth"
	"metrics/internal/handlers/collector"
	"net/http"
)

func main() {
	db := database.MemStorage{}
	h := collector.CollectorHandler{DB: &db}

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, auth.Handle)
	mux.HandleFunc(`/update/`, h.Handle)

	err := http.ListenAndServe(`:8080`, mux)

	if err != nil {
		panic(err)
	}
}
