package articles

import (
	"gin-blog/apis/users"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title        string          `gorm:"size:32" json:"title,omitempty"`                 // 标题
	Abstract     string          `json:"abstract,omitempty"`                             // 简介
	Content      string          `json:"content,omitempty"`                              // 内容
	LookCount    int             `json:"lookCount,omitempty"`                            // 浏览量
	CollectCount int             `json:"collectCount,omitempty"`                         // 收藏量
	Tags         []TagModel      `gorm:"many2many:article_tag" json:"tags,omitempty"`    //标签
	Comments     []CommentModel  `gorm:"foreignKey:ArticleID" json:"comments,omitempty"` // 评论列表
	User         users.UserModel `gorm:"foreignKey:UserID" json:"user"`                  // 文章作者
	Category     string          `json:"category,omitempty"`                             // 文章分类
	Source       string          `json:"source,omitempty"`                               // 文章来源
	Word         int             `json:"word,omitempty"`                                 //文章字数
	Banner       ImageModel      `gorm:"foreignKey:ImageID" json:"banner"`               // 文章封面
}
type CommentModel struct {
	gorm.Model      `json:"-"`
	SubComments     []*CommentModel `json:"sub_comments,omitempty" gorm:"foreignKey:ParentCommentID"`   // 子评论
	ParentComment   *CommentModel   `json:"parent_comment,omitempty" gorm:"foreignKey:ParentCommentID"` // 父集评论
	ParentCommentID *uint           `json:"parentCommentID"`                                            // 子评论ID
	Content         string          `json:"content,omitempty"`                                          // 评论内容
	Article         ArticleModel    `json:"article,omitempty" gorm:"foreignKey:ArticleID"`              // 关联的文章
	User            users.UserModel `json:"user,omitempty"`                                             // 关联的用户
}

type ImageModel struct {
	gorm.Model `json:"-"`
	Path       string `json:"path,omitempty"`                // 图片路径
	Hash       string `json:"hash,omitempty" gorm:"size:38"` //判断图片是否重复
	Name       string `json:"name,omitempty"`                // 文章名字
}

type TagModel struct {
	gorm.Model
	Title    string         `gorm:"size:16" json:"title,omitempty"`
	Articles []ArticleModel `gorm:"many2many2:article_tag" json:"-"`
}
