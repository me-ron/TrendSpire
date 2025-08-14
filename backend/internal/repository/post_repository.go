package repository

import (
	"backend/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	GetByID(id uuid.UUID) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
	Update(post *entity.Post) error
	Delete(id uuid.UUID) error
}

type PostRepositoryGorm struct {
	db *gorm.DB
}

func NewPostRepositoryGorm(db *gorm.DB) *PostRepositoryGorm {
	return &PostRepositoryGorm{db: db}
}

func (r *PostRepositoryGorm) Create(post *entity.Post) error {
	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}
	return r.db.Create(post).Error
}

func (r *PostRepositoryGorm) GetByID(id uuid.UUID) (*entity.Post, error) {
	var post entity.Post
	err := r.db.
		Preload("Likes").
		Preload("Comments").
		First(&post, "id = ?", id).Error
	return &post, err
}

func (r *PostRepositoryGorm) GetAll() ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.
		Preload("Likes").
		Preload("Comments").
		Find(&posts).Error
	return posts, err
}

func (r *PostRepositoryGorm) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepositoryGorm) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Post{}, "id = ?", id).Error
}
