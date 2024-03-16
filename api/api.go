package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	DraftApi
	TagApi
	UploadApi
}

var ApiGroupApp = new(ApiGroup)
