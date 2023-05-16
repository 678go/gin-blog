package system

import (
	"gin-blog/apis/images"
	"github.com/gin-gonic/gin"
)

func SysImagesRouter(r *gin.RouterGroup) {
	v1 := r.Group("api/v1")
	v1.POST("/images", images.ImageUpload)
}
