package routes

import (
	"mafiasu_ws/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}
	}
}
