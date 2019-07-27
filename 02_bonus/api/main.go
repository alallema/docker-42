package main

import (
	"log"
	"encoding/json"
	"net/http"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
