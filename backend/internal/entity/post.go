package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes    []Like    `gorm:"foreignKey:PostID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Like struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"` 
	PostID    uuid.UUID `gorm:"type:uuid;not null;index"` 
	CreatedAt time.Time
}

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	PostID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
