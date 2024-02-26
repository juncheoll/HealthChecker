package handler

import (
	"bytes"
	"encoding/json"
	"health-checker/data"
	"io"
	"log"
	"net/http"
)

func postGPUInfoToManager() {

	postData, err := json.Marshal(data.MyGPUInfo)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://"+*data.ManagerAddress+"gpuinfo", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server Response:", string(body))
}
