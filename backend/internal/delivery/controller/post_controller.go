package controller

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController struct {
	postUsecase usecase.PostUsecase
}

func NewPostController(postUsecase usecase.PostUsecase) *PostController {
	return &PostController{postUsecase: postUsecase}
}

// CreatePost (Authenticated)
func (pc *PostController) CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user claims from context
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(jwt.UserClaims)

	post := &entity.Post{
		AuthorID: userClaims.ID, // Always use logged-in user as author
		Title:    input.Title,
		Content:  input.Content,
	}

	if err := pc.postUsecase.CreatePost(c.Request.Context(), post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetPostByID (Public)
func (pc *PostController) GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	post, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetAllPosts (Public)
func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postUsecase.GetAllPosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// UpdatePost (Authenticated + Ownership check)
func (pc *PostController) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract claims
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(jwt.UserClaims)

	// Fetch existing post
	existingPost, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	// Ownership check
	if existingPost.AuthorID != userClaims.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to update this post"})
		return
	}

	// Update fields
	existingPost.Title = input.Title
	existingPost.Content = input.Content

	if err := pc.postUsecase.UpdatePost(c.Request.Context(), existingPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}

// DeletePost (Authenticated + Ownership check)
func (pc *PostController) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	// Extract claims
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(jwt.UserClaims)

	// Fetch existing post
	existingPost, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	// Ownership check
	if existingPost.AuthorID != userClaims.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to delete this post"})
		return
	}

	if err := pc.postUsecase.DeletePost(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}
