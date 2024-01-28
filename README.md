## 项目简介
:dart: 练习go，web项目，快速练习，前端使用开源后台管理 [vue-vben-admin](https://github.com/vbenjs/vue-vben-admin)

项目基础代码学习于 [《Go 语言编程之旅》](https://github.com/go-programming-tour-book/)

项目架构学习于 [project-layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

项目api接口对接 [vue-vben-admin](https://github.com/vbenjs/vue-vben-admin)

### 目录结构
```
├─api                 接口
├─cmd                 命令行
├─configs             配置文件
├─docs                文档集合
├─global              全局变量
├─initialize          初始化
├─internal            内部模块
│  ├─dao              数据访问层
│  ├─middleware       中间件
│  ├─model            模型层
│  ├─routers          路由层
│  └─service          服务层 - 项目核心业务逻辑
├─pkg                 模块包
├─sql                 项目sql文件
└─storage             项目生成的临时文件
    ├─logs            日志
    └─uploads         上传文件

main.go               程序入口				
```

## 链路追踪

> 需要安装 Jaeger
>
> https://www.jaegertracing.io/

<br>

> windows 使用 Jaeger
>
>1. 解压下载的 Jaeger 压缩包
>2. cd 到解压的目录，执行 jaeger-all-in-one.exe
>3. 运行 Jaeger Web UI http://localhost:16686/

<br>

`internal/routers/router.go 中间件 注入X-Trace-ID、X-Span-ID`
```go
// 链路追踪
r.Use(middleware.Tracing())
```

`pkg/logger/logger.go 日志追踪`
```go
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}
```

`internal/service/service.go sql追踪`
```go
// svc.dao = dao.New(global.DBEngine)
svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
```

## Swagger
Swagger文档 https://pkg.go.dev/github.com/swaggo/gin-swagger

> 配置cmd swag命令
>
> cd xxx\mod\github.com\swaggo\swag@v1.16.2\cmd\swag 即go get安装依赖swaggo的目录
>
> go build -> swag.exe -> 丢到go的 bin目录下

### 生成文档
```bash
swag init --help

swag init
```
api文档 http://localhost:8000/swagger/index.html