package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SuccessHealthResponse struct {
	Status          string `json:"status"`
	Message         string `json:"message"`
	Database_status string `json:"database_status"`
}

type ErrorHealthResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// HealthCheck handles the health check API and checks database connection
// @Summary      Check server health
// @Description  Returns the health status of the service and database
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  SuccessHealthResponse
// @Failure      500  {object}  ErrorHealthResponse
// @Router       /health [get]
func HealthCheck(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check database connection
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Unhealthy",
				"message": "Database connection failed",
			})
			return
		}

		// Return health status
		c.JSON(http.StatusOK, gin.H{
			"status":          "Healthy",
			"message":         "Service is running",
			"database_status": "Connected",
		})
	}
}
