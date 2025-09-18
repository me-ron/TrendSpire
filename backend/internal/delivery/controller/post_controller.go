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

type PostController struct {
	postUsecase usecase.PostUsecase
}

func NewPostController(postUsecase usecase.PostUsecase) *PostController {
	return &PostController{postUsecase: postUsecase}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var input struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		ImageURL string `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input", []response.APIError{
			{Code: "INVALID_INPUT", Detail: err.Error()},
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

	post := &entity.Post{
		AuthorID: userID,
		Title:    input.Title,
		Content:  input.Content,
		ImageURL: input.ImageURL,
	}
	if err := pc.postUsecase.CreatePost(c.Request.Context(), post); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create post", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusCreated, "Post created", post)
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	post, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Post not found", []response.APIError{
			{Code: "NOT_FOUND", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Post retrieved", post)
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postUsecase.GetAllPosts(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch posts", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Posts retrieved", posts)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "id", Code: "INVALID_UUID", Detail: err.Error()},
		})
		return
	}

	var input struct {
		Title    *string `json:"title"`
		Content  *string `json:"content"`
		ImageURL *string `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input", []response.APIError{
			{Code: "INVALID_INPUT", Detail: err.Error()},
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

	post, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Post not found", []response.APIError{
			{Code: "NOT_FOUND", Detail: err.Error()},
		})
		return
	}
	if post.AuthorID != userID {
		response.Error(c, http.StatusForbidden, "You are not allowed to update this post", []response.APIError{
			{Code: "FORBIDDEN", Detail: "You can only update your own posts"},
		})
		return
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Content != nil {
		post.Content = *input.Content
	}
	if input.ImageURL != nil {
		post.ImageURL = *input.ImageURL
	}

	if err := pc.postUsecase.UpdatePost(c.Request.Context(), post); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update post", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Post updated", post)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid post ID", []response.APIError{
			{Field: "id", Code: "INVALID_UUID", Detail: err.Error()},
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

	post, err := pc.postUsecase.GetPostByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Post not found", []response.APIError{
			{Code: "NOT_FOUND", Detail: err.Error()},
		})
		return
	}
	if post.AuthorID != userID {
		response.Error(c, http.StatusForbidden, "You are not allowed to delete this post", []response.APIError{
			{Code: "FORBIDDEN", Detail: "You can only delete your own posts"},
		})
		return
	}

	if err := pc.postUsecase.DeletePost(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete post", []response.APIError{
			{Code: "DB_ERROR", Detail: err.Error()},
		})
		return
	}
	response.Success(c, http.StatusOK, "Post deleted", nil)
}
