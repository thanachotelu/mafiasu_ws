package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck handles the health check API and checks database connection
// @Summary      Check server health
// @Description  Returns the health status of the service and database
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
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
