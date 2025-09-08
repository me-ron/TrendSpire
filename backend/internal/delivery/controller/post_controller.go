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

func (pc *PostController) CreatePost(c *gin.Context) {
	type CreatePostInput struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ImageURL string `json:"image_url"`
	}
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, err := uuid.Parse(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	post := &entity.Post{
		AuthorID: userID,
		Title:    input.Title,
		Content:  input.Content,
		ImageURL: input.ImageURL,
	}

	if err := pc.postUsecase.CreatePost(c.Request.Context(), post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}


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

func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postUsecase.GetAllPosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	type UpdatePostInput struct {
		Title    *string `json:"title"`
		Content  *string `json:"content"`
		ImageURL *string `json:"image_url"`
	}
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, err := uuid.Parse(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	existingPost, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	if existingPost.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to update this post"})
		return
	}

	if input.Title != nil {
		existingPost.Title = *input.Title
	}
	if input.Content != nil {
		existingPost.Content = *input.Content
	}
	if input.ImageURL != nil {
		existingPost.ImageURL = *input.ImageURL
	}

	if err := pc.postUsecase.UpdatePost(c.Request.Context(), existingPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}


func (pc *PostController) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(*jwt.Claims)
	userID, err := uuid.Parse(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	existingPost, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	if existingPost.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to delete this post"})
		return
	}

	if err := pc.postUsecase.DeletePost(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}
