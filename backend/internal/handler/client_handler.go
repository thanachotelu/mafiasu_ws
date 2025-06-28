package handler

import (
	"log"
	"mafiasu_ws/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService interfaces.ClientService
}

func NewClientHandler(clientService interfaces.ClientService) *ClientHandler {
	return &ClientHandler{clientService}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiKey, err := h.clientService.CreateClient(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		log.Println("❌ CreateClient error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"api_key": apiKey,
	})
}

func (h *ClientHandler) RevokeClient(c *gin.Context) {
	var req struct {
		APIKey string `json:"api_key" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.clientService.RevokeClient(c.Request.Context(), req.APIKey)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid API key"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "API key revoked",
	})
}

func (h *ClientHandler) GetUserLogs(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
		return
	}

	logs, err := h.clientService.GetLogs(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}
