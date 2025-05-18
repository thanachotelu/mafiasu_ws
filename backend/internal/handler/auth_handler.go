package handler

import (
	"mafiasu_ws/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService     interfaces.UserService
	keycloakService interfaces.KeycloakService
}

func NewAuthHandler(userService interfaces.UserService, keycloakService interfaces.KeycloakService) *AuthHandler {
	return &AuthHandler{
		userService:     userService,
		keycloakService: keycloakService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.keycloakService.Login(c, loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "Login successful",
	})
}
