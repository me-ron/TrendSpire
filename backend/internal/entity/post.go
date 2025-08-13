package entity

import (
    "time"
)

type Post struct {
	ID        uint     `gorm:"primaryKey"`
	Author    UserInfo `gorm:"embedded"`
	Title     string   `gorm:"not null"`
	Content   string   `gorm:"type:text"`
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     []Like
	Comments  []Comment
}

type Like struct {
	ID        uint     `gorm:"primaryKey"`
	User      UserInfo `gorm:"embedded"`
	PostID    uint     `gorm:"index;not null"`
	CreatedAt time.Time
}

type Comment struct {
	ID        uint     `gorm:"primaryKey"`
	User      UserInfo `gorm:"embedded"`
	PostID    uint     `gorm:"index;not null"`
	Content   string   `gorm:"type:text;not null"`
	CreatedAt time.Time
}
