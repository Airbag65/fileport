package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func verifyToken(r *http.Request) (string, error) {
	request, err := http.NewRequest("GET", "http://127.0.0.1:8000/validate", &bytes.Buffer{})
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", r.Header.Get("Authorization"))
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", &InvalidTokenError{}
	}
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return "InvalidTokenError"
}
