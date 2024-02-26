package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"health-checker/data"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	//TODO::usageStats 갱신 후 Health Check Manager에게 전달
	data.MyGPUInfo.UsageStats = getUsageStats()

	postGPUInfoToManager()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"success"})
}

func getUsageStats() data.UsageStats {
	usageStats := data.UsageStats{}

	client, err := api.NewClient(api.Config{
		Address: data.PrometheusAddress,
	})
	if err != nil {
		log.Fatalf("Error creating Prometheus client: %v\n", err)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `nv_gpu_utilization`
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		log.Fatalf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}

	fmt.Printf("Result:\n%v\n", result)

	return usageStats
}
