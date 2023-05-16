package system

import (
	"gin-blog/apis/menus"
	"github.com/gin-gonic/gin"
)

func SysMenuRouter(r *gin.Engine) {
	v1 := r.Group("api/v1")
	v1.POST("/images", menus.MenuCreateView)
}
