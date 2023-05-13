package routers

import (
	"gin-blog/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() (r *gin.Engine) {
	r = gin.New()

	r.Use(config.GinLogger(), config.GinRecovery(true))
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	return
}
