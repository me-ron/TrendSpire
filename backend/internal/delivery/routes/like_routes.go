package routes

import (
	"backend/internal/delivery/controller"
	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.RouterGroup, likeController *controller.LikeController, authMiddleware gin.HandlerFunc) {
	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.POST("/:postId/like", likeController.ToggleLike)    
		auth.GET("/:postId/likes", likeController.GetLikesByPost) 
	}
}
