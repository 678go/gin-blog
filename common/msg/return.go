package msg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// OK 成功数据
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.data = data
	if msg != "" {
		res.msg = msg
	}
	// todo msg id 用于问题定位
	//res.RequestId = tools.GenerateMsgIDFromContext(c)
	c.JSON(http.StatusOK, res.ReturnOK())
}

// Error 错误的数据
func Error(c *gin.Context, code int, msg string) {
	res := &Response{
		code: code,
		msg:  msg,
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}
