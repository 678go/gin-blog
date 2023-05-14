package msg

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
)

// OK 成功数据
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	// todo Msg id 用于问题定位
	//res.RequestId = tools.GenerateMsgIDFromContext(c)
	c.JSON(http.StatusOK, res.ReturnOK())
}

// Error 错误的数据 入参err为空会panic
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	slog.Error(msg, "Msg", err)
	c.JSON(http.StatusOK, res.ReturnError(code))
}
