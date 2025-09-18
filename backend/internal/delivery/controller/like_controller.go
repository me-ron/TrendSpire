package controller

import (
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"backend/pkg/response"
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
	postID, err := uuid.Parse(c.Param("post_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "post_id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", []response.APIError{
			{Code: "UNAUTHORIZED", Detail: "Missing or invalid authentication token"},
		})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, _ := uuid.Parse(userClaims.UserID)

	if err := lc.likeUC.ToggleLike(c.Request.Context(), postID, userID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to toggle like", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Like toggled", nil)
}

func (lc *LikeController) GetLikesByPost(c *gin.Context) {
	postID, err := uuid.Parse(c.Param("post_id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "post_id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	likes, err := lc.likeUC.GetLikesByPost(c.Request.Context(), postID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch likes", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Likes retrieved", likes)
}
