package router

import (
	"net/http"

	//_ "project-layout/docs"

	"github.com/binbinly/gin-pkg/app"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter Load loads the middlewares, routes, handlers.
func NewRouter() *gin.Engine {
	g := gin.New()

	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// HealthCheck 健康检查路由
	g.GET("/health", app.HealthCheck)
	// 静态资源，主要是图片
	g.StaticFS("/group1", http.Dir("data"))

	if app.IsLocal() {
		// swagger api docs
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// pprof router 性能分析路由
		// 默认关闭，开发环境下可以打开
		// 访问方式: HOST/debug/pprof
		// 通过 HOST/debug/pprof/profile 生成profile
		// 查看分析图 go tool pprof -http=:5000 profile
		// see: https://github.com/gin-contrib/pprof
		pprof.Register(g)
	}

	return g
}
