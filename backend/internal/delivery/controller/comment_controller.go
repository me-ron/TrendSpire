package controller

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct {
	uc usecase.CommentUsecase
}

func NewCommentController(uc usecase.CommentUsecase) *CommentController {
	return &CommentController{uc: uc}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	claims, exists := ctx.Get("user")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized", []response.APIError{
			{Code: "UNAUTHORIZED", Detail: "Missing or invalid authentication token"},
		})
		return
	}
	userClaims := claims.(*jwt.Claims)

	postID, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "post_id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input", []response.APIError{
			{Field: "content", Code: "INVALID_INPUT", Detail: err.Error()},
		})
		return
	}

	userID, _ := uuid.Parse(userClaims.UserID)
	comment := entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: req.Content,
	}

	if err := c.uc.CreateComment(ctx, &comment); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create comment", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(ctx, http.StatusCreated, "Comment created", comment)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid comment ID", []response.APIError{
			{Field: "id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	claims, exists := ctx.Get("user")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized", []response.APIError{
			{Code: "UNAUTHORIZED", Detail: "Missing or invalid authentication token"},
		})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, _ := uuid.Parse(userClaims.UserID)

	existing, err := c.uc.GetCommentByID(ctx, id)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Comment not found", []response.APIError{
			{Field: "id", Code: "NOT_FOUND", Detail: err.Error()},
		})
		return
	}
	if existing.UserID != userID {
		response.Error(ctx, http.StatusForbidden, "Forbidden", []response.APIError{
			{Code: "FORBIDDEN", Detail: "You can only update your own comments"},
		})
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input", []response.APIError{
			{Field: "content", Code: "INVALID_INPUT", Detail: err.Error()},
		})
		return
	}

	existing.Content = req.Content
	if err := c.uc.UpdateComment(ctx, existing); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update comment", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(ctx, http.StatusOK, "Comment updated", existing)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid comment ID", []response.APIError{
			{Field: "id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	claims, exists := ctx.Get("user")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized", []response.APIError{
			{Code: "UNAUTHORIZED", Detail: "Missing or invalid authentication token"},
		})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, _ := uuid.Parse(userClaims.UserID)

	existing, err := c.uc.GetCommentByID(ctx, id)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Comment not found", []response.APIError{
			{Field: "id", Code: "NOT_FOUND", Detail: err.Error()},
		})
		return
	}
	if existing.UserID != userID {
		response.Error(ctx, http.StatusForbidden, "Forbidden", []response.APIError{
			{Code: "FORBIDDEN", Detail: "You can only delete your own comments"},
		})
		return
	}

	if err := c.uc.DeleteComment(ctx, id); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to delete comment", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(ctx, http.StatusOK, "Comment deleted", nil)
}

func (c *CommentController) GetCommentsByPostID(ctx *gin.Context) {
	postID, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "post_id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	comments, err := c.uc.GetCommentsByPostID(ctx, postID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch comments", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(ctx, http.StatusOK, "Comments retrieved", comments)
}
