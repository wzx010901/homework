package models

type Comment struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string     `gorm:"size:500;not null" json:"content"`
	UserID    uint       `gorm:"not null;index" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PostID    uint       `gorm:"not null;index" json:"post_id"`
	Post      Post       `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CreatedAt CustomTime `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt CustomTime `gorm:"autoUpdateTime" json:"updated_at"`
}
