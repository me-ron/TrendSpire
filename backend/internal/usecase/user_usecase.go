package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/pkg/hash"
	"backend/pkg/jwt"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User, rawPassword string) error
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	GetProfile(ctx context.Context, id string) (*entity.User, error)
}

type userUsecase struct {
	userRepo    repository.UserRepository
	jwtService  jwt.JWTService
	timeout     time.Duration
}

func NewUserUsecase(userRepo repository.UserRepository, jwtService jwt.JWTService, timeout time.Duration) UserUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		jwtService: jwtService,
		timeout:    timeout,
	}
}

type LoginResponse struct {
	User          *entity.User `json:"user"`
	AccessToken   string       `json:"access_token"`
	RefreshToken  string       `json:"refresh_token"`
}

func (uc *userUsecase) Register(ctx context.Context, user *entity.User, rawPassword string) error {
	hashedPassword, err := hash.HashPassword(rawPassword)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	return uc.userRepo.Create(user)
}

func (uc *userUsecase) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !hash.CheckPassword(user.PasswordHash, password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := uc.jwtService.GenerateAccessToken(user.ID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := uc.jwtService.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *userUsecase) GetProfile(ctx context.Context, id string) (*entity.User, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return uc.userRepo.FindByID(ID)
}
