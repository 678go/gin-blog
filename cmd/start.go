package cmd

import (
	"context"
	"fmt"
	"gin-blog/config"
	"gin-blog/routers"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := config.InitConf(filePath); err != nil {
			return err
		}
		app := config.InitApp()

		r := routers.Setup()
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%s", app.Port),
			Handler: r,
		}

		go func() {
			// 开启一个goroutine启动服务
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				app.L.Error("监听端口失败", "msg", err)
			}
		}()

		quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
		/*
			// kill 默认会发送 syscall.SIGTERM 信号
			// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
			// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
			// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
		*/
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
		<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
		app.L.Info("关闭服务 ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
		if err := srv.Shutdown(ctx); err != nil {
			app.L.Error("关闭服务失败", "msg", err)
		}

		app.L.Info("Server exiting")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
