package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"context"

	"github.com/google/uuid"
)

type LikeUsecase interface {
	ToggleLike(ctx context.Context, postID, userID uuid.UUID) error
	GetLikesByPost(ctx context.Context, postID uuid.UUID) ([]entity.Like, error)
}

type likeUsecase struct {
	likeRepo repository.LikeRepository
}

func NewLikeUsecase(likeRepo repository.LikeRepository) LikeUsecase {
	return &likeUsecase{likeRepo: likeRepo}
}

func (uc *likeUsecase) ToggleLike(ctx context.Context, postID, userID uuid.UUID) error {
	like := &entity.Like{
		PostID: postID,
		UserID: userID,
	}
	return uc.likeRepo.ToggleLike(ctx, like)
}

func (uc *likeUsecase) GetLikesByPost(ctx context.Context, postID uuid.UUID) ([]entity.Like, error) {
	return uc.likeRepo.FindByPostID(ctx, postID)
}
