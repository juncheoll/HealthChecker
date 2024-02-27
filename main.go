package main

import (
	"flag"
	"fmt"
	"health-checker/data"
	"health-checker/handler"
	"log"
	"net/http"
)

func main() {

	handler.GetUsageStats()

	http.HandleFunc("/model-info", handler.ModelInfoHandler)
	http.HandleFunc("/prometeus", handler.PrometheusHandler)

	data.ManagerAddress = flag.String("manager", "IP:PORT", "address of manager")
	flag.Parse()

	fmt.Printf("Starting server at %s\n", data.HttpAddress)
	if err := http.ListenAndServe(data.HttpAddress, nil); err != nil {
		log.Fatal(err)
	}
}
