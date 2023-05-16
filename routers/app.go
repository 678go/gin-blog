package routers

import (
	_ "gin-blog/docs"
	"gin-blog/middleware"
	"gin-blog/routers/system"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Setup() (r *gin.Engine) {

	r = gin.New()
	// 注册gin自定义日志
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := r.Group("")

	system.SysInfoRouter(g)   // 系统管理
	system.SysImagesRouter(g) // 图片路由
	system.SysMenuRouter(r)   // 菜单路由
	return
}
