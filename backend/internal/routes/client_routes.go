package routes

import (
	"mafiasu_ws/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterClientRoutes(r *gin.Engine, clientHandler *handler.ClientHandler) {
	v1 := r.Group("/api/v1")
	{
		client := v1.Group("/client")
		{
			client.POST("/create", clientHandler.CreateClient)
			client.POST("/revoke", clientHandler.RevokeClient)
		}
	}
}
