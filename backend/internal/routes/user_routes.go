 package routes

 import (
	"mafiasu_ws/internal/handler"

 	"github.com/gin-gonic/gin"
 )

func RegisterUserRoutes(r *gin.Engine, userHandlers *handler.UserHandler) {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/:id", userHandlers.GetUserByID)
			users.GET("", userHandlers.GetAllUsers)
			users.POST("", userHandlers.AddUser)
			users.PUT("/:id", userHandlers.UpdateUser)
			users.DELETE("/:id", userHandlers.DeleteUser)
		}
	}
}
