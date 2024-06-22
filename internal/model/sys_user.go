package model

type User struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Id         uint32 `json:"id"`
	CreateTime string `json:"create_time"`
	Avatar     string `json:"avatar"`
	IsPrivacy  uint8  `json:"is_privacy"`
}

type UserInfo struct {
	Id        uint32 `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	IsPrivacy uint8  `json:"is_privacy"`
}

type BindedUserInfo struct {
	UserInfo
	Sort uint8 `json:"sort"`
}

type UserBinding struct {
	Id        uint32 `json:"id"`
	UserId1   uint32 `gorm:"column:user_id_1" json:"user_id_1"`
	UserId2   uint32 `gorm:"column:user_id_2" json:"user_id_2"`
	CreatedAt string `json:"created_at"`
	Sort      uint8  `json:"sort"`
}

type SaveBindedUserSort struct {
	Id   uint32 `json:"id" binding:"required"`
	Sort uint8  `json:"sort" binding:"required"`
}
type UserSetting struct {
	Id           uint32 `json:"id"`
	UserId       uint32 `json:"user_id"`
	PrimaryColor string `json:"primary_color"`
}

func (u User) TableName() string {
	return "user"
}
