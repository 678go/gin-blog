package users

import (
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type Role int

const (
	PermissionAdmin       Role = iota + 1 // 管理员
	PermissionUser                        // 普通用户
	PermissionVisitor                     // 游客
	PermissionDisableUser                 // 被禁用的用户
)

func (r Role) MarshalJSON() ([]byte, error) {
	var str string
	switch r {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "有错用户"
	default:
		str = "未知用户"
	}
	return json.Marshal(str)
}

type User struct {
	gorm.Model
	NickName string    `json:"nick_name,omitempty" gorm:"size:36"`                                              // 用户昵称
	UserName string    `json:"user_name,omitempty" gorm:"size:36"`                                              // 用户名
	Password string    `json:"password,omitempty" gorm:"size:64"`                                               // 密码
	Avatar   string    `json:"avatar_id,omitempty" gorm:"size:256"`                                             // 头像url
	Email    string    `json:"email,omitempty" gorm:"size:128"`                                                 // 邮箱
	Phone    string    `json:"phone,omitempty" gorm:"size:18"`                                                  // 电话
	ThereId  string    `json:"there_id,omitempty" gorm:"size:60"`                                               // 三方id
	Role     int       `json:"role,omitempty" gorm:"size:4;default:1"`                                          // 用户权限
	Articles []Article `json:"-" gorm:"foreignKey:UserId"`                                                      // 文章表
	Collects []Collect `json:"-" gorm:"many2many:user_collects;joinForeignKey:UserId;JoinReferences:ArticleId"` // 收藏了哪些文章
}

type Article struct {
}
type Collect struct {
}
