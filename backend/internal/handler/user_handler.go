package handler

import (
	"log"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) AddUser(c *gin.Context) {
    var req models.CreateUserRequest

    // ตรวจสอบว่า JSON ที่ส่งมาถูกต้องหรือไม่
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Invalid request body: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    log.Printf("Request data: %+v", req)

    // เรียกใช้ AddUser จาก userService
    user, err := h.userService.AddUser(c.Request.Context(), req)
    if err != nil {
        log.Printf("Error adding user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ส่ง Response กลับไปยัง Client
    log.Printf("User added successfully: %+v", user)
    c.JSON(http.StatusCreated, gin.H{
        "message": "User added successfully",
        "user":    user,
    })
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	updatedUser, err := h.userService.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		if strings.Contains(err.Error(), "already in use") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deletedUser, err := h.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedUser)
}