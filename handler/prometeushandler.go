package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"health-checker/data"
)

func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	//TODO::usageStats 갱신 후 Health Check Manager에게 전달
	data.MyGPUInfo.UsageStats = GetUsageStats()

	postGPUInfoToManager()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"success"})
}

func GetUsageStats() data.UsageStats {
	usageStats := data.UsageStats{}

	resp, err := http.Get(data.PrometheusAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	usageStats.Tmp = string(body)

	defer fmt.Printf("Result:\n%v\n", usageStats)

	return usageStats
}
