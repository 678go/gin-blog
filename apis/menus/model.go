package menus

import (
	"gin-blog/apis/articles"
	"gorm.io/gorm"
)

type MenuModel struct {
	gorm.Model
	Title        string                `json:"title,omitempty"`
	Abstract     string                `json:"abstract,omitempty"`                                                                            // 菜单简介
	AbstractTime int                   `json:"abstractTime,omitempty"`                                                                        // 简介切换时间
	MenuImages   []articles.ImageModel `json:"menuImages,omitempty" gorm:"many2many:menu_image;joinForeignKey:MenuID;joinForeignKey:ImageID"` // 菜单图片
	MenuTime     int                   `json:"menuTime,omitempty"`                                                                            // 菜单图片切换时间
	Sort         int                   `json:"sort,omitempty"`                                                                                // 菜单的顺序
}

// MenuImageModel 自定义菜单和背景图的连接
type MenuImageModel struct {
	Menu   MenuModel           `gorm:"foreignKey:MenuID"`
	Images articles.ImageModel `gorm:"foreignKey:ImageID"`
	Sort   int
}
