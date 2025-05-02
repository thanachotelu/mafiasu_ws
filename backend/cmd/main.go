 package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"mafiasu_ws/config"
	"mafiasu_ws/database"

	extRepo "mafiasu_ws/external/repository"
	extRoutes "mafiasu_ws/external/routes"
	extService "mafiasu_ws/external/services"
	extInterfaces "mafiasu_ws/external/interfaces"
	intHandler "mafiasu_ws/internal/handler"
	intRepo "mafiasu_ws/internal/repository"
	intRoutes "mafiasu_ws/internal/routes"
	intService "mafiasu_ws/internal/service"
)

func main() {
	// โหลด configuration
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
	kc := extService.NewKeycloakService(cfg.Keycloak)
	userService := intService.NewUserService(userRepo,kc)
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
	waitForKeycloak(cfg.Keycloak.BaseURL)
	initializeRoles(kc)
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
func initializeRoles(keycloakService extInterfaces.KeycloakService) {
    roles := []string{"user", "Affiliator"} // รายชื่อ Role ที่ต้องการสร้าง
    for _, role := range roles {
        if err := keycloakService.CreateRoleIfNotExists(role); err != nil {
            log.Printf("Failed to create role '%s': %v", role, err)
        }
    }
}
func waitForKeycloak(baseURL string) {
    for {
        resp, err := http.Get(baseURL)
        if err == nil && resp.StatusCode == http.StatusOK {
            log.Println("Keycloak is ready")
            return
        }
        log.Println("Waiting for Keycloak...")
        time.Sleep(5 * time.Second)
    }
}