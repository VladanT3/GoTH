package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/VladanT3/GoTH/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.GetLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started on http://localhost" + listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
