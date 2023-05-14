package routers

import (
	"gin-blog/middleware"
	"gin-blog/routers/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() (r *gin.Engine) {

	r = gin.New()
	// 注册gin自定义日志
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	g := r.Group("")

	// 系统管理
	system.SysInfoRouter(g)

	// 图片路由
	system.SysImagesRouter(g)
	return
}
