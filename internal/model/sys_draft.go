package model

type Draft struct {
	Id         uint32 `json:"id"`
	UserId     uint32 `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	IsPublish  uint8  `json:"is_publish"`
	IsDelete   uint8  `json:"is_delete"`
}
