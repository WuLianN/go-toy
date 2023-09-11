package model

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Id uint `json:"id"`
}

type Role struct {
	// UserId uint `json:"user_id"`
	Name string `json:"name"`
	Value string `json:"value"`
}

func (u User) TableName() string {
	return "user"
}