package model

import (
	"github.com/golang-jwt/jwt/v4"
	"navigator/api"
)

type model struct {
}

type BaseResponse struct {
	*api.Code
	Data interface{} `json:"data"`
}

const (
	EnvProduction = "production"
)

const (
	StatusSuccess = 2
	StatusFail    = 1
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}
