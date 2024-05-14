package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["status"] = "ok"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func hostname(w http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("<h1>VERSION 3: Hello from Health Checker %v</h1>", name)))
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/hostname", hostname)
	http.ListenAndServe(":8000", nil)
}
