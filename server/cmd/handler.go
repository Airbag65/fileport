package main

import (
	"log/slog"
	"net/http"
)

func listDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	email, err := verifyToken(r)
	if err != nil {
		slog.Error("verify token error", "err", err)
		Unauthorized(w)
		return
	}

	dir, err := GetUserDirPath(email)
	if err != nil {
		InternalServerError(w)
		return
	}

	WriteJSON(w, dir)
}
