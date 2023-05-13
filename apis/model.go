package apis

import (
	"gin-blog/common"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName string      `json:"nick_name,omitempty" gorm:"size:36"`                                              // 用户昵称
	UserName string      `json:"user_name,omitempty" gorm:"size:36"`                                              // 用户名
	Password string      `json:"password,omitempty" gorm:"size:64"`                                               // 密码
	Avatar   string      `json:"avatar_id,omitempty" gorm:"size:256"`                                             // 头像url
	Email    string      `json:"email,omitempty" gorm:"size:128"`                                                 // 邮箱
	Phone    string      `json:"phone,omitempty" gorm:"size:18"`                                                  // 电话
	ThereId  string      `json:"there_id,omitempty" gorm:"size:60"`                                               // 三方id
	Role     common.Role `json:"role,omitempty" gorm:"size:4;default:1"`                                          // 用户权限
	Articles []Article   `gorm:"foreignKey:UserId"`                                                               // 文章表
	Collects []Article   `json:"-" gorm:"many2many:user_collects;joinForeignKey:UserId;JoinReferences:ArticleID"` // 收藏了哪些文章
}

type Article struct {
	gorm.Model
	Title        string    `gorm:"size:32" json:"title,omitempty"`   // 标题
	Abstract     string    `json:"abstract,omitempty"`               // 简介
	Content      string    `json:"content,omitempty"`                // 内容
	LookCount    int       `json:"lookCount,omitempty"`              // 浏览量
	CollectCount int       `json:"collectCount,omitempty"`           // 收藏量
	Tags         []Tag     `gorm:"many2many:article_tags" json:"-"`  //标签
	Comments     []Comment `gorm:"foreignKey:ArticleId" json:"-"`    // 评论列表
	User         User      `gorm:"foreignKey:UserId"`                // 文章作者
	UserId       uint      `json:"user_id"`                          // 用户ID
	Category     string    `json:"category,omitempty"`               // 文章分类
	Source       string    `json:"source,omitempty"`                 // 文章来源
	Word         int       `json:"word,omitempty"`                   //文章字数
	Image        Image     `gorm:"foreignKey:ImageId" json:"banner"` // 文章封面
	ImageId      uint      `json:"image_id"`
}

type Comment struct {
	gorm.Model
	SubComments     []*Comment `json:"sub_comments,omitempty" gorm:"foreignKey:ParentCommentID"`   // 子评论
	ParentComment   *Comment   `json:"parent_comment,omitempty" gorm:"foreignKey:ParentCommentID"` // 父集评论
	ParentCommentID *uint      `json:"parentCommentID"`                                            // 子评论ID
	Content         string     `json:"content,omitempty"`                                          // 评论内容
	Article         Article    `json:"-" gorm:"foreignKey:ArticleId"`                              // 关联的文章
	ArticleId       uint       `json:"article_id"`                                                 // 文章id
	User            User       `json:"user,omitempty"`                                             // 关联的用户
	UserId          uint       `json:"user_id"`                                                    // 用户ID
}

type Image struct {
	gorm.Model
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash" gorm:"size:38"` //判断图片是否重复
	Name string `json:"name"`                // 文章名字
}

type Tag struct {
	gorm.Model
	Title    string    `gorm:"size:16" json:"title,omitempty"`
	Articles []Article `gorm:"many2many:article_tags;" json:"-"`
}
