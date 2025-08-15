package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username      string    `gorm:"uniqueIndex;not null"`
	Email         string    `gorm:"uniqueIndex;not null"`
	PasswordHash  string    `gorm:"not null"`
	ProfilePicURL string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Posts         []Post `gorm:"foreignKey:AuthorID;references:ID"`
}


type UserInfo struct {
	ID            uuid.UUID `gorm:"type:uuid;not null"`
	Username      string    `gorm:"not null"`
	ProfilePicURL string
}
