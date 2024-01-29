package model

type User struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Id         uint32 `json:"id"`
	CreateTime string `json:"create_time"`
}

func (u User) TableName() string {
	return "user"
}
