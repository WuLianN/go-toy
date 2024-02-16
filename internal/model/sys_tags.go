package model

type Tag struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	UserId  uint32 `json:"user_id"`
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
}

type DraftTag struct {
	TagId   uint32 `json:"tag_id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
	DraftId uint32 `json:"draft_id"`
}

type DraftTags struct {
	Tags    []Tag  `json:"tags"`
	DraftId uint32 `json:"draft_id"`
}

type DraftWithTags struct {
	Tags []Tag `json:"tags" gorm:"foreignKey:Id"`
	Draft
}
