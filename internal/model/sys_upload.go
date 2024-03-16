package model

type UploadRecord struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	AccessUrl string `json:"access_url"`
}
