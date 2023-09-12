package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
}

var ApiGroupApp = new(ApiGroup)
