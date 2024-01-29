package routers

import (
	"net/http"
	"time"

	docs "github.com/WuLianN/go-toy/docs"
	"github.com/WuLianN/go-toy/global"
	"github.com/WuLianN/go-toy/internal/middleware"
	"github.com/WuLianN/go-toy/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",     // 自定义键值对名称
		FillInterval: time.Second, // 间隔多久时间放 Quantum 个令牌
		Capacity:     10,          // 令牌桶的容量
		Quantum:      10,          // 每次到达间隔时间后所放的具体令牌数量
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

	// 系统基础组
	systemBaseGroup := r.Group("/")
	{
		InitBaseRouter(systemBaseGroup)
		InitMenuRouter(systemBaseGroup)
	}

	// 系统权限组
	systemAuthGroup := r.Group("/")
	systemAuthGroup.Use(middleware.JWT())
	{
		InitUserRouter(systemAuthGroup)
		InitStatisticsRouter(systemAuthGroup)
	}

	return r
}
