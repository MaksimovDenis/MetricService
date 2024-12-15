package api

import (
	"fmt"
	"net/http"
	"strings"
	"yandexCourse/metricService/service/internal/storage"
)

func GetMerics(ms *storage.MemStorage, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) != 4 || parts[1] != "update" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if parts[2] == "" || parts[3] == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		strConv, err := storage.StringConverter(parts[3])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ms.AddMetrics(parts[2], strConv); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

		fmt.Printf("Метрика %v со значенем %v успешно обнавлена\n", parts[2], strConv)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
