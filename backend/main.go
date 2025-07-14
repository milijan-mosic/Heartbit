package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HelloResponse{
		Message: "Hello, World!",
		Status:  http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to encode JSON"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	address := ":5000"
	urlPrefix := "/api/1.0"

	http.HandleFunc(urlPrefix+"/", indexHandler)

	log.Printf("Server starting on: http://localhost%s\n", address)
	http.ListenAndServe(address, nil)
}
