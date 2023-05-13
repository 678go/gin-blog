package common

import "github.com/goccy/go-json"

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
