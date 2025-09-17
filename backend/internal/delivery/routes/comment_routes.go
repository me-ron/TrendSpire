package routes

import (
	"backend/internal/delivery/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup, commentController *controller.CommentController, authMiddleware gin.HandlerFunc) {
	r.GET("/post/:postID", commentController.GetCommentsByPostID)

	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.POST("/", commentController.CreateComment)
		auth.PUT("/:id", commentController.UpdateComment)
		auth.DELETE("/:id", commentController.DeleteComment)
	}
}
