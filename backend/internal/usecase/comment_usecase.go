package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"context"

	"github.com/google/uuid"
)

type CommentUsecase interface {
	CreateComment(ctx context.Context, comment *entity.Comment) error
	UpdateComment(ctx context.Context, comment *entity.Comment) error
	DeleteComment(ctx context.Context, commentID uuid.UUID) error
	GetCommentsByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Comment, error)
}

type commentUsecase struct {
	repo repository.CommentRepository
}

func NewCommentUsecase(repo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{repo: repo}
}

func (uc *commentUsecase) CreateComment(ctx context.Context, comment *entity.Comment) error {
	return uc.repo.Create(ctx, comment)
}

func (uc *commentUsecase) UpdateComment(ctx context.Context, comment *entity.Comment) error {
	return uc.repo.Update(ctx, comment)
}

func (uc *commentUsecase) DeleteComment(ctx context.Context, commentID uuid.UUID) error {
	return uc.repo.Delete(ctx, commentID)
}

func (uc *commentUsecase) GetCommentsByPostID(ctx context.Context, postID uuid.UUID) ([]entity.Comment, error) {
	return uc.repo.FindByPostID(ctx, postID)
}
