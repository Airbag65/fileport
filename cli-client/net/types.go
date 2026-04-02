package net

import "net/http"

var (
	client = &http.Client{}
)

type ValidateTokenReq struct {
	AuthToken string `json:"auth_token"`
	Email     string `json:"email"`
}

type ResponseCode int

const (
	OK                  ResponseCode = 200
	BadRequset          ResponseCode = 400
	NotFound            ResponseCode = 404
	Unauthorized        ResponseCode = 401
	NotModified         ResponseCode = 304
	ImATeapot           ResponseCode = 418
	InternalServerError ResponseCode = 500
)
