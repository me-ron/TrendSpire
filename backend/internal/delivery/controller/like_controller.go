package controller

import (
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LikeController struct {
	likeUC usecase.LikeUsecase
}

func NewLikeController(likeUC usecase.LikeUsecase) *LikeController {
	return &LikeController{likeUC: likeUC}
}

func (lc *LikeController) ToggleLike(c *gin.Context) {
	postIDParam := c.Param("post_id")
	postID, err := uuid.Parse(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userClaims := claims.(*jwt.Claims)

	userID, err := uuid.Parse(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := lc.likeUC.ToggleLike(c.Request.Context(), postID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to toggle like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "like toggled successfully"})
}

func (lc *LikeController) GetLikesByPost(c *gin.Context) {
	postIDParam := c.Param("post_id")
	postID, err := uuid.Parse(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	likes, err := lc.likeUC.GetLikesByPost(c.Request.Context(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch likes"})
		return
	}

	c.JSON(http.StatusOK, likes)
}
