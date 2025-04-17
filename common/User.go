package common

import (
	"gorm.io/gorm"
)

// User 用户模型，符合 GORM 规范
type User struct {
	gorm.Model        // 嵌入 GORM 默认字段（ID, CreatedAt, UpdatedAt, DeletedAt）
	Username   string `gorm:"unique;not null"` // 用户名，唯一且不能为空
	Mobile     string `gorm:"unique;not null"` // 手机号，唯一且不能为空
	Avatar     string `gorm:"default:null"`    // 头像链接，默认为空
	Password   string `gorm:"not null"`        // 加密后的密码，不能为空
}
