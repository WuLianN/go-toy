package api

type ApiGroup struct {
	BaseApi
	AuthApi
}

var ApiGroupApp = new(ApiGroup)
