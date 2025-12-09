package error

import "fmt"

var (
	ErrUserAlreadyExists = &AppError{Code: 1001, Message: "用户已存在"}
	ErrRoleNotFound      = &AppError{Code: 1002, Message: "角色不存在"}
	ErrInvalidPermission = &AppError{Code: 1003, Message: "权限无效"}
	// ... 其他业务错误
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
