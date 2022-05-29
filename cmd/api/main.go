package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type Config struct {
	port int
	env  string
}
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.port, "port", 4000, "port the server listens on")
	flag.StringVar(&cfg.env, "env", "dev", "application environment: dev|pro")
	flag.Parse()
	fmt.Printf("server running on port %d with application environment: %s", cfg.port, cfg.env)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}
		js, _ := json.MarshalIndent(currentStatus, "", "\t")
		w.Header().Set("content-type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
