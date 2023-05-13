package system

import "gorm.io/gorm"

type SysModel struct {
	gorm.Model
	Title string `json:"title"`
}
