package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("connection to /home")
	WriteJSON(w, "OK")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("connection to /auth/login")
	if !ensureJSON(w, r) {
		return
	}
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		BadRequest(w)
		return
	}
	loginReq.Password = encryptPassword(loginReq.Password)
	authportReq := NewAuthportLoginRequest(loginReq.Email, loginReq.Password, loginReq.ClientIdentifier, GetRequestIP(r))

	// FIXME: Send to authport and check response
	WriteJSON(w, authportReq)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("connection to /auth/new")
	if !ensureJSON(w, r) {
		return
	}
	var regReq RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		BadRequest(w)
		return
	}
	regReq.Password = encryptPassword(regReq.Password)
	// FIXME: Send to authport and check response

	WriteJSON(w, regReq)
}
