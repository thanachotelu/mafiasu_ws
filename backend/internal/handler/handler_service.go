package handler

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
)

type UserHandler struct {
	userService interfaces.UserService
}

// NewUserHandler constructs a new UserHandler with the given UserService
func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserByID handles GET /users/:id requests
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// AddUser handles POST /users requests
func (h *UserHandler) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// เรียกใช้ RegisterUser จาก service
	err := h.userService.RegisterUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
