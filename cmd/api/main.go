package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
type application struct {
	config Config
	logger *log.Logger
}

func main() {
	var cfg Config
	logger :=
		log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	flag.IntVar(&cfg.port, "port", 4000, "port the server listens on")
	flag.StringVar(&cfg.env, "env", "dev", "application environment: dev|pro")
	flag.Parse()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}
	logger.Println("starting server on port ", cfg.port)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Println(err)
	}

}
