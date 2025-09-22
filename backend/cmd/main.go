package main

import (
	"backend/config"
	"backend/internal/delivery/controller"
	"backend/internal/delivery/routes"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)
	redisClient := config.InitRedis(cfg)
	cloud := config.InitCloudinary(cfg)

	log.Println("DB:", db, "Redis:", redisClient, "Cloudinary:", cloud)

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepositoryGorm(db)
	likeRepo := repository.NewLikeRepositoryGorm(db)
	commentRepo := repository.NewCommentRepositoryGorm(db)

	jwtService := jwt.NewJWTService(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		"TrendSpire",
	)

	userUC := usecase.NewUserUsecase(userRepo, jwtService, 5*time.Second)
	postUC := usecase.NewPostUsecase(postRepo)
	likeUC := usecase.NewLikeUsecase(likeRepo)
	commentUC := usecase.NewCommentUsecase(commentRepo)

	userController := controller.NewUserController(userUC)
	postController := controller.NewPostController(postUC)
	likeController := controller.NewLikeController(likeUC)
	commentController := controller.NewCommentController(commentUC)

	authMiddleware := middleware.AuthMiddleware(jwtService)

	r := gin.Default()
	routes.SetupRoutes(r, userController, postController, likeController, commentController, authMiddleware)

	r.Run(":" + cfg.App.Port)
}
