package api

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	StatisticsApi
	DraftApi
	TagApi
}

var ApiGroupApp = new(ApiGroup)
