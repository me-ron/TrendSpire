package repository

import (
	"context"
	"errors"

	"backend/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeRepository interface {
	FindByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Like, error)
	Exists(ctx context.Context, postID, userID uuid.UUID) (bool, error)
	ToggleLike(ctx context.Context, like *entity.Like) error
}

type likeRepositoryGorm struct {
	db *gorm.DB
}

func NewLikeRepositoryGorm(db *gorm.DB) LikeRepository {
	return &likeRepositoryGorm{db: db}
}

func (r *likeRepositoryGorm) FindByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Like, error) {
	var likes []entity.Like
	err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&likes).Error
	return likes, err
}

func (r *likeRepositoryGorm) Exists(ctx context.Context, postID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Like{}).
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *likeRepositoryGorm) ToggleLike(ctx context.Context, like *entity.Like) error {
	var existing entity.Like
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", like.User.ID, like.PostID).
		First(&existing).Error

	if err == nil {
		return r.db.WithContext(ctx).Delete(&existing).Error
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Assign a new UUID if not set
	if like.ID == uuid.Nil {
		like.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(like).Error
}
