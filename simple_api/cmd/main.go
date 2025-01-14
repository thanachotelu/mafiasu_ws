package main

import (
	_ "users/docs" // ให้ Swag สร้างเอกสารใน Folder docs โดยอัตโนมัติ
	"users/internal/middleware"

	"log"
	"users/internal/config"
	"users/internal/handler"
	"users/internal/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Simple API Example
// @version         1.0
// @description     This is a simple example of using Gin with Swagger.
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cfg := config.LoadConfig()

	db := user.InitDB("localhost", "5432", "postgres", "postgres123", "postgres")
	defer db.Close()

	r := gin.Default()

	// Swagger endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handler.HealthCheck(db))

	// User API routes
	authRequired := r.Group("/api/v1", middleware.BearerAuth(cfg.APIToken))
	{
		// @Security BearerAuth
		authRequired.PUT("/users/:id", handler.UpdateUserHandler(db))
		authRequired.DELETE("/users/:id", handler.DeleteUserHandler(db))
		authRequired.POST("/users", handler.AddUserHandler(db))
		authRequired.GET("/users", handler.GetAllUsersHandler(db))
		authRequired.GET("/users/:id", handler.GetUserByID(db))
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
