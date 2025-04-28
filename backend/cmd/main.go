package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"mafiasu_ws/config"
	"mafiasu_ws/database"

	extRepo "mafiasu_ws/external/repository"
	extRoutes "mafiasu_ws/external/routes"
	extService "mafiasu_ws/external/services"

	intHandler "mafiasu_ws/internal/handler"
	intRepo "mafiasu_ws/internal/repository"
	intRoutes "mafiasu_ws/internal/routes"
	intService "mafiasu_ws/internal/service"
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

	// User
	userRepo := intRepo.NewUserRepository(db.GetPool())
	userService := intService.NewUserService(userRepo)
	userHandler := intHandler.NewUserHandler(userService)

	// Car
	carRepo := intRepo.NewCarRepository(db.GetPool())
	carService := intService.NewCarService(carRepo)
	carHandler := intHandler.NewCarHandler(carService)

	// Booking
	bookingRepo := intRepo.NewBookingRepository(db.GetPool())
	bookingService := intService.NewBookingService(bookingRepo)
	bookingHandler := intHandler.NewBookingHandler(bookingService)

	// Affiliates
	affiliateRepo := extRepo.NewAffiliateRepository(db.GetPool())
	affiliateService := extService.NewAffiliateService(affiliateRepo)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Register routes
	intRoutes.RegisterUserRoutes(r, userHandler)
	intRoutes.RegisterCarRoutes(r, carHandler)
	intRoutes.RegisterBookingRoutes(r, bookingHandler)
	extRoutes.RegisterAffiliateRoutes(r, affiliateService)

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
