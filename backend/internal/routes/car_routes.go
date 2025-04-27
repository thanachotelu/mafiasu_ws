package routes

import (
	"mafiasu_ws/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCarRoutes(r *gin.Engine, carHandlers *handler.CarHandler) {
	v1 := r.Group("/api/v1")
	{
		cars := v1.Group("/cars")
		{
			cars.GET("/:id", carHandlers.GetCarByID)
			cars.GET("", carHandlers.GetAllCars)
			cars.POST("", carHandlers.AddCar)
			cars.PUT("/:id", carHandlers.UpdateCar)
			cars.DELETE("/:id", carHandlers.DeleteCar)
		}
	}
}
