package menus

import (
	"gin-blog/apis"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Title        string       `json:"title,omitempty"`
	Abstract     string       `json:"abstract,omitempty"`                                                                        // 菜单简介
	AbstractTime int          `json:"abstractTime,omitempty"`                                                                    // 简介切换时间
	Images       []apis.Image `json:"Images,omitempty" gorm:"many2many:menu_image;joinForeignKey:MenuID;JoinReferences:ImageID"` // 菜单图片
	MenuTime     int          `json:"menuTime,omitempty"`                                                                        // 菜单图片切换时间
	Sort         int          `json:"sort,omitempty"`                                                                            // 菜单的顺序
}

// MenuImage 自定义菜单和背景图的连接
type MenuImage struct {
	MenuId  uint
	Menu    Menu `gorm:"foreignKey:MenuId"`
	ImageId uint
	Images  apis.Image `gorm:"foreignKey:ImageId"`
	Sort    int
}
