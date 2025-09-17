package routes

import (
	"backend/internal/delivery/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userController *controller.UserController,
	postController *controller.PostController,
	likeController *controller.LikeController,
	commentController *controller.CommentController,
	authMiddleware gin.HandlerFunc,
) {
	api := router.Group("/api/v1")

	// User routes
	UserRoutes(api.Group("/users"), userController, authMiddleware)

	// Post routes
	RegisterPostRoutes(api.Group("/posts"), postController, authMiddleware)

	// Like routes
	LikeRoutes(api.Group("/likes"), likeController, authMiddleware)

	// Comment routes
	CommentRoutes(api.Group("/comments"), commentController, authMiddleware)
}
