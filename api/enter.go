package api

import (
	"github.com/WuLianN/go-toy/api/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
