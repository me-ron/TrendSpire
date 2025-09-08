package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"context"

	"github.com/google/uuid"
)

type PostUsecase interface {
	CreatePost(ctx context.Context, post *entity.Post) error
	GetPostByID(ctx context.Context, id uuid.UUID) (*entity.Post, error)
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error
}

type postUsecase struct {
	postRepo repository.PostRepository
}

func NewPostUsecase(postRepo repository.PostRepository) PostUsecase {
	return &postUsecase{postRepo: postRepo}
}

func (u *postUsecase) CreatePost(ctx context.Context, post *entity.Post) error {
	return u.postRepo.Create(post)
}

func (u *postUsecase) GetPostByID(ctx context.Context, id uuid.UUID) (*entity.Post, error) {
	return u.postRepo.GetByID(id)
}

func (u *postUsecase) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	return u.postRepo.GetAll()
}

func (u *postUsecase) UpdatePost(ctx context.Context, post *entity.Post) error {
	return u.postRepo.Update(post)
}

func (u *postUsecase) DeletePost(ctx context.Context, id uuid.UUID) error {
	return u.postRepo.Delete(id)
}
