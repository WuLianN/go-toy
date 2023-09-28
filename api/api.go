package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	StatisticsApi
}

var ApiGroupApp = new(ApiGroup)
