package main

import (
	"metrics/internal/database"
	"metrics/internal/handlers/auth"
	"metrics/internal/handlers/collect_metric"
	"net/http"
)

func main() {
	db := database.MemStorage{}
	h := collect_metric.CollectorHandler{DB: &db}

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, auth.Handle)
	mux.HandleFunc(`/update/`, h.Handle)

	err := http.ListenAndServe(`:8080`, mux)

	if err != nil {
		panic(err)
	}
}
