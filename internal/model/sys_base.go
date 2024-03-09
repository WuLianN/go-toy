package model

type ResponseResult struct {
	Code    uint32 `json:"code" example:"0"`
	Message string `json:"message" example:"ok"`
	Result  any    `json:"result"`
}

type RecommendList struct {
	Id         uint32 `json:"id"`
	UserId     uint32 `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	BgImage    string `json:"bg_image"`

	Tags []Tag `json:"tags" gorm:"foreignKey:Id"`
}
