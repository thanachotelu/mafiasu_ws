package handler

import (
	"mafiasu_ws/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct {
	authRepo interfaces.AuthRepository
}

func NewMiddlewareHandler(authRepo interfaces.AuthRepository) *MiddlewareHandler {
	return &MiddlewareHandler{authRepo}
}

func (h *MiddlewareHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API Key required"})
			c.Abort()
			return
		}

		clientID, err := h.authRepo.ValidateAPIKey(c.Request.Context(), apiKey) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			c.Abort()
			return
		}

		c.Set("client_id", clientID)
		c.Next()
	}
}

func (h *MiddlewareHandler) LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID, exists := c.Get("client_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		err := h.authRepo.LogRequest(c.Request.Context(), clientID.(int), c.Request.URL.Path, c.Request.Method) 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log request"})
			c.Abort()
			return
		}

		c.Next()
	}
}
