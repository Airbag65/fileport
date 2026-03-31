package main

import (
	"log/slog"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("connection to /home")
		WriteJSON(w, "OK")
	})

	// NOTE: Not needed! Use AUTHPORT directly in client instead
	// server.HandleFunc("POST /auth/login", loginHandler)
	// server.HandleFunc("POST /auth/new", registerHandler)

	server.HandleFunc("GET /files/list", listDirectoryHandler)

	handler := cors.Default().Handler(server)
	http.ListenAndServe(":8001", handler)
}
