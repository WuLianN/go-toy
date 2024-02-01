package model

type Tag struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type DraftTag struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	DraftId uint32 `json:"draft_id"`
}
