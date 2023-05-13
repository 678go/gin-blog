package config

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var cfg *ini.File

type App struct {
	Name string `ini:"name"`
	Port string `ini:"port"`
	L    *slog.Logger
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

func InitApp() *App {
	app := new(App)
	if err := cfg.Section("app").MapTo(app); err != nil {
		slog.Error("初始化app失败", "msg", err)
		os.Exit(-1)
	}
	app.L = slog.Default().With("model", "app")
	slog.Info("初始化app成功!")
	return app
}

func InitDB() *gorm.DB {
	var l logger.LogLevel
	m := new(Mysql)
	if err := cfg.Section("mysql").MapTo(m); err != nil {
		slog.Error("初始化数据库信息失败", "msg", err)
		os.Exit(-1)
	}
	switch m.logLevel {
	case "error":
		l = logger.Error
	default:
		l = logger.Info
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(l),
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
	sql.SetConnMaxIdleTime(1 * time.Hour)
	slog.Info("初始化数据库连接信息成功!")
	return db
}

func InitConf(path string) (err error) {
	l := new(Log)
	cfg, err = ini.Load(path)
	if err != nil {
		fmt.Println("配置文件路径错误", err)
		return err
	}
	if err := cfg.Section("log").MapTo(l); err != nil {
		fmt.Println("初始化日志失败,", err)
		return err
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
	return
}
