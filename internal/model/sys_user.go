package model

type User struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Id         uint32 `json:"id"`
	CreateTime string `json:"create_time"`
}

type UserInfo struct {
	Id       uint32 `json:"id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

func (u User) TableName() string {
	return "user"
}
