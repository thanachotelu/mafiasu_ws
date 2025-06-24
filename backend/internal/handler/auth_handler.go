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

	// ดึง user จาก database เพื่อเอา role
	user, err := h.userService.GetUserByUsername(c, loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"role":    user.Role,
		"message": "Login successful",
	})
}
