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
