package model

type Tag struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	UserId uint32 `json:"user_id"`
}

type DraftTag struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	DraftId uint32 `json:"draft_id"`
}
