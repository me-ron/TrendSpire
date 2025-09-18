package repository

import (
	"context"

	"backend/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, comment *entity.Comment) error
	Delete(ctx context.Context, commentID uuid.UUID) error
	FindByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, id uuid.UUID) (*entity.Comment, error)
}

type commentRepositoryGorm struct {
	db *gorm.DB
}

func NewCommentRepositoryGorm(db *gorm.DB) CommentRepository {
	return &commentRepositoryGorm{db: db}
}

func (r *commentRepositoryGorm) Create(ctx context.Context, comment *entity.Comment) error {
	if comment.ID == uuid.Nil {
		comment.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *commentRepositoryGorm) Update(ctx context.Context, comment *entity.Comment) error {
	return r.db.WithContext(ctx).Save(comment).Error
}

func (r *commentRepositoryGorm) Delete(ctx context.Context, commentID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Comment{}, "id = ?", commentID).Error
}

func (r *commentRepositoryGorm) FindByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.WithContext(ctx).Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

func (r *commentRepositoryGorm) GetCommentByID(ctx context.Context, id uuid.UUID) (*entity.Comment, error) {
    var comment entity.Comment
    if err := r.db.WithContext(ctx).First(&comment, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &comment, nil
}
