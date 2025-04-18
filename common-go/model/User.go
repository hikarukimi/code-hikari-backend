package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型，符合 GORM 规范
type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null;size:255"`
	Mobile       string `gorm:"unique;not null;size:15"`
	Email        string `gorm:"unique;size:100"`
	Avatar       string `gorm:"size:255"`
	Password     string `gorm:"not null;size:255"`
	Role         int16  `gorm:"not null"`
	Exp          int
	Level        int
	LastSignInAt time.Time `gorm:"type:timestamp with time zone"`
	// 定义多对多关联
	RelatedUsers []User `gorm:"many2many:user_relations;joinForeignKey:UserID;joinReferences:TargetID"`
}

// UserRelation 用户关系模型，符合 GORM 规范
type UserRelation struct {
	gorm.Model
	UserID       int64 `gorm:"not null"`
	TargetID     int64 `gorm:"not null"`
	RelationType int16 `gorm:"not null"`
}

// UserSignIn 用户签到记录模型，符合 GORM 规范
type UserSignIn struct {
	gorm.Model
	UserID     int64     `gorm:"not null"`
	SignInDate time.Time `gorm:"type:date;not null"`
}

// UserFavorite 用户收藏模型，符合 GORM 规范
type UserFavorite struct {
	gorm.Model
	UserID     int64  `gorm:"not null"`
	TargetID   int64  `gorm:"not null"`
	TargetType int16  `gorm:"not null"`
	Password   string `gorm:"not null"` // 加密后的密码，不能为空
}
