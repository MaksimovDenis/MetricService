package main

import (
	"net/http"
	"yandexCourse/metricService/service/internal/api"
	"yandexCourse/metricService/service/internal/storage"
)

func main() {
	data := make(map[string]any)
	newSorage := storage.New(data)

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		api.GetMerics(newSorage, w, r)
	})

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
