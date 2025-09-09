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
	IsPrivacy  uint8  `json:"is_privacy"`
	BgImage    string `json:"bg_image"`
}

type CreateDraftResponse struct {
	Draft_id uint32 `json:"draft_id" example:"1"`
}

type SearchDraft struct {
	UserId     uint32 `json:"user_id"`
	Keyword    string `json:"keyword"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	IsSelf     uint8  `json:"is_self"`
	SerachType uint8  `json:"serach_type"`
}
