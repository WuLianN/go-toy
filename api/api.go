package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	DraftApi
	TagApi
}

var ApiGroupApp = new(ApiGroup)
