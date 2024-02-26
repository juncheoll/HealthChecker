package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"health-checker/data"
)

func ModelInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var modelInfo data.ModelInfo
	err := json.NewDecoder(r.Body).Decode(&modelInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received data: %+v", modelInfo)

	//TODO::modelInfo 갱신 후 Health Checker Manager에게 전달
	data.MyGPUInfo.ModelInfo = modelInfo

	postGPUInfoToManager()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"success"})
}
