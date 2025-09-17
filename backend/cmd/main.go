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
	// Load config
	cfg := config.LoadConfig()

	// Init external dependencies
	db := config.InitDB(cfg)
	redisClient := config.InitRedis(cfg)
	cloud := config.InitCloudinary(cfg)

	log.Println("DB:", db, "Redis:", redisClient, "Cloudinary:", cloud)

	// Init repositories
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepositoryGorm(db)
	likeRepo := repository.NewLikeRepositoryGorm(db)
	commentRepo := repository.NewCommentRepositoryGorm(db)

	// Init services
	jwtService := jwt.NewJWTService(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		"MyApp",
	)

	// Init usecases
	userUC := usecase.NewUserUsecase(userRepo, jwtService, 5*time.Second)
	postUC := usecase.NewPostUsecase(postRepo)
	likeUC := usecase.NewLikeUsecase(likeRepo)
	commentUC := usecase.NewCommentUsecase(commentRepo)

	// Init controllers
	userController := controller.NewUserController(userUC)
	postController := controller.NewPostController(postUC)
	likeController := controller.NewLikeController(likeUC)
	commentController := controller.NewCommentController(commentUC)

	// Middleware
	authMiddleware := middleware.AuthMiddleware(jwtService)

	// Setup routes
	r := gin.Default()
	routes.SetupRoutes(r, userController, postController, likeController, commentController, authMiddleware)

	// Run server
	r.Run(":" + cfg.App.Port)
}
