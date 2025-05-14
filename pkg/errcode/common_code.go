package errcode

var (
	Success                   = NewError(0, "ok")
	Fail                      = NewError(-1, "fail")
	Warning                   = NewError(777, "注意")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	Unauthorized              = NewError(10000003, "无权限")
	UnauthorizedAuthNotExist  = NewError(10000004, "鉴权失败, 无Token")
	UnauthorizedTokenError    = NewError(10000005, "鉴权失败, Token错误")
	UnauthorizedTokenTimeout  = NewError(10000006, "鉴权失败, Token超时")
	UnauthorizedTokenGenerate = NewError(10000007, "鉴权失败, Token生成失败")
	TooManyRequests           = NewError(10000008, "请求过多")
)
