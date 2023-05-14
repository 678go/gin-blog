package cmd

import (
	"gin-blog/apis"
	"gin-blog/apis/menus"
	"gin-blog/apis/system"
	"gin-blog/config"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init",
	Long:  "init",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// 初始化各种数据表
func run() {
	if err := config.InitConf(filePath); err != nil {
		return
	}
	db := config.InitDB()
	slog.Info("开始初始化表结构:")
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	if err := db.AutoMigrate(
		&apis.Tag{},
		&system.System{},
		&apis.User{},
		&apis.Image{},
		&system.LoginData{},
		&apis.Comment{},
		&apis.Comment{},
		&menus.Menu{},
		&menus.MenuImage{},
		&apis.Article{}); err != nil {
		return
	}
	slog.Info("表结构初始化完成")
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
