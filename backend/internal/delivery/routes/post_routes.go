package routes

import (
	"backend/internal/delivery/controller"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.RouterGroup, postController *controller.PostController, authMiddleware gin.HandlerFunc) {
	r.GET("/", postController.GetAllPosts)
	r.GET("/:id", postController.GetPostByID)

	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.POST("/", postController.CreatePost)
		auth.PUT("/:id", postController.UpdatePost)
		auth.DELETE("/:id", postController.DeletePost)
	}
}
