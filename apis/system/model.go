package system

import (
	"gin-blog/apis"
	"gorm.io/gorm"
)

type System struct {
	gorm.Model
	Title string `json:"title"`
}

type LoginData struct {
	gorm.Model
	UserId uint      `json:"user_id"`
	User   apis.User `json:"-" gorm:"foreignKey:UserId"`
	IP     string    `json:"IP"`
	Device string    `json:"device"`
}
