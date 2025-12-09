package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type model struct {
}

const (
	EnvProduction = "production"
	SuccessCode   = 0
	SystemErr     = 500
	StatusSuccess = 2
	StatusFail    = 1
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}
