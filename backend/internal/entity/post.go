package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Author    UserInfo  `gorm:"embedded"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     []Like
	Comments  []Comment
}

type Like struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	User      UserInfo  `gorm:"embedded"`
	PostID    uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
}

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	User      UserInfo  `gorm:"embedded"`
	PostID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
