package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	StatisticsApi
	DraftApi
}

var ApiGroupApp = new(ApiGroup)
