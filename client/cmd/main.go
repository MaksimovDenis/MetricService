package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

var (
	PollCount   int64
	RandomValue float64

	pollInterval = time.Second * 3

	address = "http://localhost:8080/"
)

func collectMetrics() map[string]any {
	metricsStorage := make(map[string]any)
	var metrics runtime.MemStats

	runtime.ReadMemStats(&metrics)

	metricsStorage["Alloc"] = (float64(metrics.Alloc))
	metricsStorage["BuckHashSys"] = float64(metrics.BuckHashSys)
	metricsStorage["Frees"] = float64(metrics.Frees)
	metricsStorage["GCCPUFraction"] = float64(metrics.GCCPUFraction)
	metricsStorage["GCSys"] = float64(metrics.GCSys)
	metricsStorage["HeapAlloc"] = float64(metrics.HeapAlloc)
	metricsStorage["HeapIdle"] = float64(metrics.HeapIdle)
	metricsStorage["HeapInuse"] = float64(metrics.HeapInuse)
	metricsStorage["HeapObjects"] = float64(metrics.HeapObjects)
	metricsStorage["HeapReleased"] = float64(metrics.HeapReleased)
	metricsStorage["HeapSys"] = float64(metrics.HeapSys)
	metricsStorage["LastGC"] = float64(metrics.LastGC)
	metricsStorage["Lookups"] = float64(metrics.Lookups)
	metricsStorage["MCacheInuse"] = float64(metrics.MCacheInuse)
	metricsStorage["MCacheSys"] = float64(metrics.MCacheSys)
	metricsStorage["MSpanInuse"] = float64(metrics.MSpanInuse)
	metricsStorage["MSpanSys"] = float64(metrics.MSpanSys)
	metricsStorage["Mallocs"] = float64(metrics.Mallocs)
	metricsStorage["NextGC"] = float64(metrics.NextGC)
	metricsStorage["NumForcedGC"] = float64(metrics.NumForcedGC)
	metricsStorage["NumGC"] = float64(metrics.NumGC)
	metricsStorage["OtherSys"] = float64(metrics.OtherSys)
	metricsStorage["PauseTotalNs"] = float64(metrics.PauseTotalNs)
	metricsStorage["StackInuse"] = float64(metrics.StackInuse)
	metricsStorage["StackSys"] = float64(metrics.StackSys)
	metricsStorage["Sys"] = float64(metrics.Sys)
	metricsStorage["TotalAlloc"] = float64(metrics.TotalAlloc)

	PollCount++
	RandomValue = rand.Float64() * 10

	metricsStorage["PollCount"] = PollCount
	metricsStorage["RandomValue"] = RandomValue

	return metricsStorage
}

func PostMetrics(metrics map[string]any) {
	url := address + "update"

	for key, value := range metrics {
		urlWithParams := fmt.Sprintf("%s/%s/%v", url, key, value)

		resp, err := http.Post(urlWithParams, "text/plain", nil)
		if err != nil {
			fmt.Println("Ошибка при отправке POST-запроса:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Ошибка при обновлении метрик, статус:", resp.Status)
		} else {
			fmt.Printf("Запрос метрики %v со значением %v успешно отправлен\n", key, value)
		}
	}
}

func main() {
	metrics := collectMetrics()

	ticker := time.NewTicker(pollInterval)

	for {
		select {
		case <-ticker.C:
			PostMetrics(metrics)
		}
	}
}
