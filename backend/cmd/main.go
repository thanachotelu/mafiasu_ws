package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"mafiasu_ws/config"
	"mafiasu_ws/database"
	"mafiasu_ws/internal/handler"
	"mafiasu_ws/internal/repository"
	"mafiasu_ws/internal/routes"
	"mafiasu_ws/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	//User
	userRepo := repository.NewUserRepository(db.GetPool())
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//Car
	carRepo := repository.NewCarRepository(db.GetPool())
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)

	//Booking
	bookingRepo := repository.NewBookingRepository(db.GetPool())
	bookingService := service.NewBookingService(bookingRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
	routes.RegisterUserRoutes(r, userHandler)
	routes.RegisterCarRoutes(r, carHandler)
	routes.RegisterBookingRoutes(r, bookingHandler)

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
