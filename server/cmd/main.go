package main

import (
	"fmt"
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
	server.HandleFunc("GET /files/get", getFileHandler)
	server.HandleFunc("POST /files/upload", uploadFileHandler)
	server.HandleFunc("POST /files/mkdir", mkdirHandler)
	server.HandleFunc("DELETE /files/delete", deleteHandler)
	server.HandleFunc("DELETE /files/rmdir", rmdirHandler)
	server.HandleFunc("PUT /files/move", moveHandler)

	handler := cors.Default().Handler(server)
	fmt.Println("Listening on :8001")
	http.ListenAndServe(":8001", handler)
}
