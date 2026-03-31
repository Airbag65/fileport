package main

type LoginRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	ClientIdentifier string `json:"client_identifier"`
}

type AuthportLoginRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	ClientIdentifier string `json:"client_identifier"`
	RemoteAddr       string `json:"remote_addr"`
}

func NewAuthportLoginRequest(email, password, clientIdentifier, remoteAddr string) *AuthportLoginRequest {
	return &AuthportLoginRequest{
		Email:            email,
		Password:         password,
		ClientIdentifier: clientIdentifier,
		RemoteAddr:       remoteAddr,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type LoginResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	AuthToken       string `json:"auth_token"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
}

type User struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Id            int    `json:"id"`
	AuthToken     string `json:"auth_token"`
	LoggedInCount int    `json:"logged_in_count"`
}
