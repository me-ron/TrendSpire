package routes

import (
	"backend/internal/delivery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userController *controller.UserController,
	postController *controller.PostController,
	authMiddleware gin.HandlerFunc,
) {
	api := router.Group("/api/v1")

	// User routes
	UserRoutes(api.Group("/users"), userController, authMiddleware)

	// Post routes
	RegisterPostRoutes(api.Group("/posts"), postController, authMiddleware)
}
