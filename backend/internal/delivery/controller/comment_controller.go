package controller

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"backend/pkg/jwt"
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
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    userClaims := claims.(*jwt.Claims)

    // Extract post ID from path
    postIDParam := ctx.Param("post_id")
    postID, err := uuid.Parse(postIDParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    var req struct {
        Content string `json:"content" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    userID, err := uuid.Parse(userClaims.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    comment := entity.Comment{
        UserID:  userID,
        PostID:  postID,
        Content: req.Content,
    }

    if err := c.uc.CreateComment(ctx, &comment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, comment)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
        return
    }

    claims, exists := ctx.Get("user")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    userClaims := claims.(*jwt.Claims)
    userID, err := uuid.Parse(userClaims.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    existingComment, err := c.uc.GetCommentByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

    if existingComment.UserID != userID {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own comments"})
        return
    }

    var req struct {
        Content string `json:"content" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    existingComment.Content = req.Content

    if err := c.uc.UpdateComment(ctx, existingComment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, existingComment)
}


func (c *CommentController) DeleteComment(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
        return
    }

    claims, exists := ctx.Get("user")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    userClaims := claims.(*jwt.Claims)
    userID, err := uuid.Parse(userClaims.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    existing, err := c.uc.GetCommentByID(ctx, id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
        return
    }

    if existing.UserID != userID {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own comments"})
        return
    }

    if err := c.uc.DeleteComment(ctx, id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}


func (c *CommentController) GetCommentsByPostID(ctx *gin.Context) {
	postID, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := c.uc.GetCommentsByPostID(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}
