package repository

import (
	"context"

	"backend/internal/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, commentID uint) error
	FindByPostID(ctx context.Context, postID uint) ([]entity.Comment, error)
}

type commentRepositoryGorm struct {
	db *gorm.DB
}

func NewCommentRepositoryGorm(db *gorm.DB) CommentRepository {
	return &commentRepositoryGorm{db: db}
}

func (r *commentRepositoryGorm) Create(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *commentRepositoryGorm) Update(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Save(comment).Error
}

func (r *commentRepositoryGorm) Delete(ctx context.Context, commentID uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Comment{}, commentID).Error
}

func (r *commentRepositoryGorm) FindByPostID(ctx context.Context, postID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}