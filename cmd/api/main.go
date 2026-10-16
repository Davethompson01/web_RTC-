package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	log := log.New(os.Stdout, "WebRtc Server: ", log.LstdFlags)

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("Failed to load .env")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Unable to start server %v", err)
		}
	}()

	// signchannel

}
