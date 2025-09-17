package routes

import (
	"backend/internal/delivery/controller"
	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.RouterGroup, likeController *controller.LikeController, authMiddleware gin.HandlerFunc) {
	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.POST("/:post_id", likeController.ToggleLike)    
		auth.GET("/:post_id", likeController.GetLikesByPost) 
	}
}
