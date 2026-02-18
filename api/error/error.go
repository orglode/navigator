package error

import (
	"fmt"
)

var (
	// ======================base code =========================//
	Success          = &AppError{Code: 0, Message: "success"}
	MissingParameter = &AppError{Code: 499, Message: "缺少参数"}
	InternalError    = &AppError{Code: 500, Message: "网络异常请稍后重试"}

	// ======================base code =========================//
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"error"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}
