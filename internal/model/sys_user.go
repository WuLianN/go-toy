package model

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u User) TableName() string {
	return "user"
}