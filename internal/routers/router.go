package routers

import (
	"net/http"
	"time"

	"go-toy/global"
	"go-toy/internal/middleware"
	"go-toy/internal/routers/api"
	test "go-toy/internal/routers/api/test"
	auth "go-toy/internal/routers/api/auth"
	"go-toy/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "go-toy/docs"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth", // 自定义键值对名称
		FillInterval: time.Second, // 间隔多久时间放 Quantum 个令牌
		Capacity:     10, // 令牌桶的容量
		Quantum:      10, // 每次到达间隔时间后所放的具体令牌数量
	},
)


func SetupRouter() *gin.Engine {
  r := gin.Default()

	// 跨域
	r.Use(middleware.Cors())

	// 访问日志
	r.Use(middleware.AccessLog())
	// 链路追踪
	r.Use(middleware.Tracing())
	// 接口限流控制
	r.Use(middleware.RateLimiter(methodLimiters))
	// 统一超时管理
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
  
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 静态资源
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 上传文件
	upload := api.NewUpload()
  r.POST("/upload/file", upload.UploadFile)
  
	// 测试组
	apiTest := r.Group("api/test")
	{
		apiTest.GET("/ping", test.Ping)
	}

	// 权限组
	apiAuth := r.Group("api/auth") 
	apiAuth.Use(middleware.JWT())
	{	
		apiAuth.GET("/checkAuth", auth.CheckAuth)
	}
	
  return r
}