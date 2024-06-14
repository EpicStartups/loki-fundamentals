package main

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	// Configure slog to log to a file (same path as carnivorous-garden)
	logPath := "./gologs/app.log"
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelError,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	// Simulate continuous logging with random numbers
	go func() {
		for {
			randomNumber := rand.Intn(100) // Generate a random number between 0 and 99
			slog.Info("Random number generated from slog", slog.Int("random_number", randomNumber))
			slog.Warn("Random number generated from slog", slog.Int("random_number", randomNumber))
			slog.Debug("Random number generated from slog", slog.Int("random_number", randomNumber))
			slog.Error("Random number generated from slog", slog.Int("random_number", randomNumber))
			time.Sleep(5 * time.Second) // Log every 5 seconds
		}
	}()

	// Create a health route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
