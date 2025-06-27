package middleware

import (
	"log"
	"mafiasu_ws/internal/interfaces"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct {
	authRepo interfaces.AuthRepository
}

func NewMiddlewareHandler(authRepo interfaces.AuthRepository) *MiddlewareHandler {
	return &MiddlewareHandler{authRepo}
}

func (h *MiddlewareHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") && !strings.HasPrefix(authHeader, "APIKey ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}
		// ====== validate token ======
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if len(tokenStr) < 10 {
				log.Println("Token string too short: ", tokenStr)
			}
			claims, err := h.authRepo.ValidateJWTToken(c.Request.Context(), tokenStr)
			if err != nil {
				log.Println("Error validating token: ", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}

			username := claims["preferred_username"]
			userID := claims["sub"]
			roles := extractRoles(claims)

			c.Set("username", username)
			c.Set("user_id", userID)
			c.Set("roles", roles)
			c.Next()
			return
		}

		// ===== validate API Key =====
		if strings.HasPrefix(authHeader, "APIKey ") {
			apiKey := strings.TrimPrefix(authHeader, "APIKey ")
			clientID, err := h.authRepo.ValidateAPIKey(c.Request.Context(), apiKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
				c.Abort()
				return
			}
			c.Set("client_id", clientID)

			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
		c.Abort()
	}
}

func (h *MiddlewareHandler) LogMiddleware() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	userID, exists := c.Get("user_id")
	// 	if !exists {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	err := h.authRepo.LogRequest(c.Request.Context(), userID.(string), c.Request.URL.Path, c.Request.Method)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log request"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	c.Next()
	// }

	return func(c *gin.Context) {
		var clientID *int
		var userID *string

		if cidRaw, ok := c.Get("client_id"); ok {
			cid := cidRaw.(int)
			clientID = &cid
		}

		if uidRaw, ok := c.Get("user_id"); ok {
			uid := uidRaw.(string)
			userID = &uid
		}

		err := h.authRepo.LogRequest(
			c.Request.Context(),
			clientID,
			userID,
			c.Request.URL.Path,
			c.Request.Method,
		)
		if err != nil {
			log.Printf("Could not log request for endpoint %s: %v", c.Request.URL.Path, err)
		}

		c.Next()
	}
}

func extractRoles(claims map[string]interface{}) []string {
	var roles []string
	if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
		if r, ok := realmAccess["roles"].([]interface{}); ok {
			for _, role := range r {
				if roleStr, ok := role.(string); ok {
					roles = append(roles, roleStr)
				}
			}
		}
	}
	return roles
}
