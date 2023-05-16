package images

import (
	"fmt"
	"gin-blog/common/msg"
	"github.com/gin-gonic/gin"
	"path"
)

// ImageUpload 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Accept json
// @Produce json
// @Success 200 {object} msg.Response{}
// @Router /images [post]
func ImageUpload(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		msg.Error(c, -1, err, "上传文件失败")
		return
	}
	fileList, ok := multipartForm.File["images"]
	if !ok {
		msg.Error(c, -1, fmt.Errorf("文件不存在"), "文件不存在")
		return
	}
	for _, f := range fileList {
		join := path.Join("test", f.Filename)
		if err := c.SaveUploadedFile(f, join); err != nil {
			msg.Error(c, -1, err, "保存文件失败")
			continue
		}
	}
	msg.OK(c, "", "上传成功")
}
