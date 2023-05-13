package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SysInfoRouter(r *gin.RouterGroup) {
	r.GET("/info", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
}
