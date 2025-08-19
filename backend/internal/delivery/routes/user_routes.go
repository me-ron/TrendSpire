package routes

import (
	"backend/internal/delivery/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController *controller.UserController, authMiddleware gin.HandlerFunc) {

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)


	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.GET("/profile", userController.GetProfile)
	}
}
