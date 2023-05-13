package config

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var l = slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("model", "gin")

// GinLogger 接收gin的默认日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		l.Info(path,
			slog.Int("status", c.Writer.Status()),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("query", query),
			slog.String("ip", c.ClientIP()),
			slog.String("user-agent", c.Request.UserAgent()),
			slog.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			slog.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					l.Error(c.Request.URL.Path,
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					l.Error("[Recovery from panic]",
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
						slog.String("stack", string(debug.Stack())),
					)
				} else {
					l.Error("[Recovery from panic]",
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
