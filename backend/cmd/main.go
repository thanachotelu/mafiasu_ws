package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mafiasu_ws/config"
	"mafiasu_ws/database"
	_ "mafiasu_ws/docs" // This will be generated
	"mafiasu_ws/internal/handler"
	"mafiasu_ws/internal/repository"
	"mafiasu_ws/internal/routes"
	"mafiasu_ws/internal/service"
)

// @title          MafiaCar API
// @version        1.0
// @description    APIs for MafiaCar affiliate user program
// @host          localhost:8000
// @BasePath      /api/v1

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

	// Enable CORS (if needed)
	// r.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
	// 	c.Next()
	// })

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
