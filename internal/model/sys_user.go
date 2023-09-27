package model

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Id uint `json:"id"`
	CreateTime string `json:"create_time"`
}

type Role struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

func (u User) TableName() string {
	return "user"
}