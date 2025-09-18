package controller

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"backend/pkg/jwt"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUC usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUC}
}

func (uc *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
	}

	if err := uc.userUsecase.Register(c.Request.Context(), user, req.Password); err != nil {
		response.Error(c, http.StatusConflict, "Failed to register user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	resp, err := uc.userUsecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	response.Success(c, http.StatusOK, "Login successful", resp)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	claimsValue, exists := c.Get("user")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	claims, ok := claimsValue.(*jwt.Claims)
	if !ok {
		response.Error(c, http.StatusUnauthorized, "Invalid token claims", nil)
		return
	}

	user, err := uc.userUsecase.GetProfile(c.Request.Context(), claims.UserID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", nil)
		return
	}

	response.Success(c, http.StatusOK, "Profile retrieved", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"created":  user.CreatedAt,
	})
}
