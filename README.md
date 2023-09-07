项目架构借鉴 https://github.com/golang-standards/project-layout/blob/master/README_zh.md


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

### Jaeger报错, 不运行链路追踪启动项目, 需要注释相关代码

`cmd/main.go 初始化Jaeger`
```go
// err = setupTracer()
// if err != nil {
// 	log.Fatalf("init.setupTracer err: %v", err)
// }
```

`pkg/logger/logger.go 日志追踪`
```go
import (
	// "github.com/gin-gonic/gin"
)

func (l *Logger) WithTrace() *Logger {
	// ginCtx, ok := l.ctx.(*gin.Context)
	// if ok {
	// 	return l.WithFields(Fields{
	// 		"trace_id": ginCtx.MustGet("X-Trace-ID"),
	// 		"span_id":  ginCtx.MustGet("X-Span-ID"),
	// 	})
	// }
	return l
}
```

`internal/routers/router.go 中间件 注入X-Trace-ID、X-Span-ID`
```go
// 链路追踪
// r.Use(middleware.Tracing())
```

`pkg/db/db.go sql追踪`
```go
// otgorm.AddGormCallbacks(db)
```

`internal/service/service.go`
```go
import (
	// otgorm "go-toy/pkg/opentracing-gorm"
)

// svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
svc.dao = dao.New(global.DBEngine)
```


## Swagger
https://pkg.go.dev/github.com/swaggo/gin-swagger

> 配置cmd swag命令
>
> cd xxx\mod\github.com\swaggo\swag@v1.16.2\cmd\swag 即go get安装依赖swaggo的目录
>
> go build -> swag.exe -> 丢到go的 bin目录下

### 生成文档
```bash
swag init --help

swag init -d ./api -o ./docs
```
http:127.0.0.1:8000/swagger/index.html