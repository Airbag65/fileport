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

// func NewRegisterRequest(email, password, name, surname string) *RegisterRequest {
// 	return &RegisterRequest{
// 		Email:    email,
// 		Password: password,
// 		Name:     name,
// 		Surname:  surname,
// 	}
// }
