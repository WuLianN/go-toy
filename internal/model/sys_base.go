package model

type VisitInfo struct {
	VisitTime string `json:"visit_time"`
	IP        string `json:"ip"`
}

type ResponseResult struct {
	Code    uint32 `json:"code" example:"0"`
	Message string `json:"message" example:"ok"`
	Result  any    `json:"result"`
}
