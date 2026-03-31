package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /home", homeHandler)
	server.HandleFunc("POST /auth/login", loginHandler)
	server.HandleFunc("POST /auth/new", registerHandler)

	handler := cors.Default().Handler(server)
	http.ListenAndServe(":8001", handler)
}
