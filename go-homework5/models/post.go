package models

type Post struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string     `gorm:"size:200;not null" json:"title"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	UserID    uint       `gorm:"not null;index" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt CustomTime `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt CustomTime `gorm:"autoUpdateTime" json:"updated_at"`
	Comments  []Comment  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}
