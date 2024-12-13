package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	port        = ":8080"
	timezone    = "America/Sao_Paulo"
	contentType = "application/json"
)

type Response struct {
	Name        string `json:"name"`
	CurrentTime string `json:"current_time"`
}

func StartServer() {
	mux := http.NewServeMux()
	registerRoutes(mux)

	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Troubles starting the server: %v", err)
	}
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handleRoot)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	currentTime, err := getCurrentTimeInLocation(timezone)
	if err != nil {
		log.Println(err)
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	response := Response{
		Name:        "Projeto Korp",
		CurrentTime: currentTime,
	}

	w.Header().Set("Content-Type", contentType)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Response error", http.StatusInternalServerError)
	}
}

func getCurrentTimeInLocation(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("error catching timezone: %w", err)
	}
	return time.Now().In(location).Format("2006-01-02T15:04:05-0700"), nil
}
