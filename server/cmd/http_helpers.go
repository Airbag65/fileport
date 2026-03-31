package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func WriteJSON(w http.ResponseWriter, v any) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(v)
}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte("Bad request"))
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("Internal server error"))
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Write([]byte("Unauthorized"))
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("Not found"))
}

func GetRequestIP(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}

func ensureJSON(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		BadRequest(w)
		return false
	}
	return true
}
