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
	// โหลด configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// เชื่อมต่อกับฐานข้อมูล
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// สร้าง repository และ services สำหรับ User
	userRepo := repository.NewUserRepository(db.GetPool())
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)


	// สร้าง middleware สำหรับ authentication
	authRepo := repository.NewAuthRepository(db.GetPool(), cfg.KeycloakPublicKey)
	middleware := handler.NewMiddlewareHandler(authRepo)

	// สร้าง gin router
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Register routes
	routes.RegisterRoutes(r, userHandler)

	

	// ใช้ middleware สำหรับ authentication
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(), middleware.LogMiddleware())

	// เริ่มต้น server
	r.Run(":8080")
}
