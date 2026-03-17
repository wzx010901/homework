package models

import (
	"time"

	"gorm.io/gorm"
)

// DB 导出的数据库连接实例
var DB *gorm.DB

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"posts,omitempty"`
	Comments  []Comment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Username == "" {
		return gorm.ErrInvalidData
	}
	if u.Password == "" {
		return gorm.ErrInvalidData
	}
	if u.Email == "" {
		return gorm.ErrInvalidData
	}
	return nil
}
