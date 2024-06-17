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
	router.Get("/foo", handlers.Make(handlers.Foo))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started on http://localhost" + listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
