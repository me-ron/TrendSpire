package repository

import (
	"context"
	"errors"

	"backend/internal/entity"

	"gorm.io/gorm"
)

type LikeRepository interface {
	FindByPostID(ctx context.Context, postID uint) ([]entity.Like, error)
	Exists(ctx context.Context, postID, userID uint) (bool, error)
	ToggleLike(ctx context.Context, like *entity.Like) error
}

type likeRepositoryGorm struct {
	db *gorm.DB
}

func NewLikeRepositoryGorm(db *gorm.DB) LikeRepository {
	return &likeRepositoryGorm{db: db}
}

func (r *likeRepositoryGorm) FindByPostID(ctx context.Context, postID uint) ([]entity.Like, error) {
	var likes []entity.Like
	err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&likes).Error
	return likes, err
}

func (r *likeRepositoryGorm) Exists(ctx context.Context, postID, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Like{}).
		Where("post_id = ? AND id = ?", postID, userID).
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

    return r.db.WithContext(ctx).Create(like).Error 
}
