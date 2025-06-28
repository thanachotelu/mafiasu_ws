package handler

import (
	"mafiasu_ws/internal/interfaces"
	"net/http"
    "github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService     interfaces.UserService
	keycloakService interfaces.KeycloakService
}

func NewAuthHandler(userService interfaces.UserService, keycloakService interfaces.KeycloakService) *AuthHandler {
	return &AuthHandler{
		userService:     userService,
		keycloakService: keycloakService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
    var loginReq struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    token, refreshToken, err := h.keycloakService.Login(c, loginReq.Username, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    parsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token"})
        return
    }
    claims := parsed.Claims.(jwt.MapClaims)

    // ดึง user_id จาก claims["sub"]
    userID, _ := claims["sub"].(string)

    roles := []string{}
    if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
        if r, ok := realmAccess["roles"].([]interface{}); ok {
            for _, v := range r {
                if roleStr, ok := v.(string); ok {
                    roles = append(roles, roleStr)
                }
            }
        }
    }

    validRoles := map[string]bool{
        "user":       true,
        "Affiliator": true,
        "admin":      true,
    }

    filteredRoles := []string{}
    for _, role := range roles {
        if validRoles[role] {
            filteredRoles = append(filteredRoles, role)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "token":         token,
        "refresh_token": refreshToken,
        "user_id":       userID, // <-- เพิ่มตรงนี้
        "roles":         filteredRoles,
        "message":       "Login successful",
    })
}
func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var req struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    token, refreshToken, err := h.keycloakService.RefreshToken(c, req.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }
    parsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token"})
        return
    }
    claims := parsed.Claims.(jwt.MapClaims)
    userID, _ := claims["sub"].(string) // <-- เพิ่มตรงนี้

    roles := []string{}
    if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
        if r, ok := realmAccess["roles"].([]interface{}); ok {
            for _, v := range r {
                if roleStr, ok := v.(string); ok {
                    roles = append(roles, roleStr)
                }
            }
        }
    }
    validRoles := map[string]bool{
        "user":       true,
        "Affiliator": true,
        "admin":      true,
    }
    filteredRoles := []string{}
    for _, role := range roles {
        if validRoles[role] {
            filteredRoles = append(filteredRoles, role)
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "token":         token,
        "refresh_token": refreshToken,
        "user_id":       userID, // <-- เพิ่มตรงนี้
        "roles":         filteredRoles,
        "message":       "Token refreshed",
    })
}