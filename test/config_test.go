package test

import (
	"gin-blog/config"
	"github.com/goccy/go-json"
	"golang.org/x/exp/slog"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))

}

// User 用户表
type User struct {
	Id       uint
	Name     string
	Articles []Article `gorm:"foreignKey:UserId"`
}

// Article 文章列表
type Article struct {
	Id    uint
	Title string
	//	UserId uint
	User User `gorm:"foreignKey:UserId"`
}

func TestDB(t *testing.T) {
	if err := config.InitConf("../config/dev.ini"); err != nil {
		return
	}
	db := config.InitDB()
	slog.Info("开始初始化表结构:")
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	db.AutoMigrate(User{}, Article{})
}

type Status int

const (
	Running Status = 1
	Except  Status = 2
	OffLine Status = 3
)

func (s Status) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "OffLine"
	}
	return json.Marshal(str)
}
