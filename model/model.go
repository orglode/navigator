package model

import "github.com/orglode/navigator/api"

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
