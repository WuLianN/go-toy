package model

type Tag struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	UserId uint32 `json:"user_id"`
}

type DraftTag struct {
	TagId   uint32 `json:"tag_id"`
	Name    string `json:"name"`
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
