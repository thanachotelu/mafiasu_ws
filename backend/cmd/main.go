 package main

// import (
// 	"log"
// 	"net/http"

// 	"mafiasu_ws/config"
// 	"mafiasu_ws/database"
// 	"mafiasu_ws/external/keycloak"
// 	"mafiasu_ws/internal/handler"
// 	"mafiasu_ws/internal/repository"
// 	"mafiasu_ws/internal/routes"
// 	"mafiasu_ws/internal/service"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// โหลด configuration
// 	cfg, err := config.New()
// 	if err != nil {
// 		log.Fatalf("failed to load config: %v", err)
// 	}

// 	// เชื่อมต่อกับฐานข้อมูล Postgres
// 	db, err := database.NewPostgresDB(cfg)
// 	if err != nil {
// 		log.Fatalf("failed to connect to database: %v", err)
// 	}

// 	// สร้าง Keycloak service
// 	keycloakSvc := keycloak.NewKeycloakService(cfg.Keycloak)

// 	// สร้าง repository และ services สำหรับ User
// 	userRepo := repository.NewUserRepository(db.GetPool())
// 	userService := service.NewUserService(userRepo, keycloakSvc)
// 	userHandler := handler.NewUserHandler(userService)

// 	// สร้าง middleware สำหรับ authentication
// 	authRepo := repository.NewAuthRepository(db.GetPool(), cfg.KeycloakPublicKey)
// 	middleware := handler.NewMiddlewareHandler(authRepo)

// 	// สร้าง gin router
// 	r := gin.Default()
// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
// 	})

// 	// Register public routes
// 	routes.RegisterRoutes(r, userHandler)

// 	// Group ที่ต้องใช้ Authentication
// 	auth := r.Group("/")
// 	auth.Use(middleware.AuthMiddleware(), middleware.LogMiddleware())
// 	// ตัวอย่าง: routes.RegisterProtectedRoutes(auth, userHandler)

// 	// เริ่มต้น server
// 	r.Run(":8080")
// }
