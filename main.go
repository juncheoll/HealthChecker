package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"health-checker/data"
	"health-checker/handler"
)

func main() {
	http.HandleFunc("/model-info", handler.ModelInfoHandler)
	http.HandleFunc("/prometeus", handler.PrometheusHandler)

	data.ManagerAddress = flag.String("manager", "IP:PORT", "address of manager")

	fmt.Printf("Starting server at %s\n", data.HttpAddress)
	if err := http.ListenAndServe(data.HttpAddress, nil); err != nil {
		log.Fatal(err)
	}

}
