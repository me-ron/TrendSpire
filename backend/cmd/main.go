package main

import (
	"backend/config"
	"backend/internal/delivery/controller"
	"backend/internal/repository"
	"backend/internal/usecase"
	"backend/internal/delivery/routes"
	"backend/internal/middleware"
	"backend/pkg/jwt"
	"time"

	"github.com/gin-gonic/gin"
	"log"

)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)
	redisClient := config.InitRedis(cfg)
	cloud := config.InitCloudinary(cfg)

	log.Println(db, redisClient, cloud)
	userRepo := repository.NewUserRepository(db)

	jwtService := jwt.NewJWTService(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		"MyApp", 
	)
	userUC := usecase.NewUserUsecase(userRepo, jwtService, 5*time.Second)
	userController := controller.NewUserController(userUC)
	authMiddleware := middleware.AuthMiddleware(jwtService)
	r := gin.Default()
	routes.SetupRoutes(r, userController, authMiddleware)

	r.Run(":" + cfg.App.Port)

	// http
}
