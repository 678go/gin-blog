package config

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var (
	cfg  *ini.File
	path string
)

type App struct {
	name string `ini:"name"`
	port string `ini:"port"`
}

type Mysql struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	logLevel string `ini:"loglevel"`
	MaxConn  int    `ini:"maxconn"`
	MaxIdle  int    `ini:"maxidle"`
}

type Log struct {
	Level string `ini:"level"`
}

func InitApp() (app *App) {
	//app := new(App)
	if err := cfg.Section("app").MapTo(app); err != nil {
		slog.Error("初始化app失败", "msg", err)
		os.Exit(-1)
	}
	return
}

func InitDB() *gorm.DB {
	m := new(Mysql)
	if err := cfg.Section("mysql").MapTo(m); err != nil {
		slog.Error("初始化数据库信息失败", "msg", err)
		os.Exit(-1)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: false,
		},
	})
	if err != nil {
		slog.Error("初始化数据库连接失败", err)
		os.Exit(-1)
	}
	sql, _ := db.DB()
	sql.SetMaxOpenConns(m.MaxConn)
	sql.SetMaxIdleConns(m.MaxIdle)
	return db
}

func init() {
	cfg, _ = ini.Load(path)
	l := new(Log)
	if err := cfg.Section("log").MapTo(l); err != nil {
		fmt.Println("初始化日志失败,", err)
		return
	}
	opts := &slog.HandlerOptions{
		AddSource: true,
	}
	switch l.Level {
	case "info":
		opts.Level = slog.LevelInfo
	default:
		opts.Level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, opts)))
	slog.Info("初始化日志成功!")
	slog.With("model", "init")
}
