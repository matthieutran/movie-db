package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const version = "1.0.0"

type StatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currStatus := StatusResponse{
			Status:  "Available",
			Version: version,
		}

		js, err := json.MarshalIndent(currStatus, "", "\t")
		if err != nil {
			fmt.Println("Could not marshal status:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Could not serve HTTP server:", err)
	}
}
