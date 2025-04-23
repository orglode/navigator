package model

type User struct {
	UserId      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
