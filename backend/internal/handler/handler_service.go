 package handler

// import (
	
// 	"net/http"
//     "mafiasu_ws/external/keycloak"
// 	"github.com/gin-gonic/gin"
// 	"mafiasu_ws/internal/interfaces"
// 	"mafiasu_ws/internal/models"
    
// )

// type UserHandler struct {
// 	userService interfaces.UserService
//     keycloakService interfaces.KeycloakService

// }

// // NewUserHandler constructs a new UserHandler with the given UserService
// func NewUserHandler(userService interfaces.UserService) *UserHandler {
// 	return &UserHandler{userService: userService}
// }

// // GetUserByID handles GET /users/:id requests
// func (h *UserHandler) GetUserByID(c *gin.Context) {
// 	id := c.Param("id")
// 	user, err := h.userService.GetUserByID(c.Request.Context(), id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// // AddUser handles POST /users requests
// func (h *UserHandler) AddUser(c *gin.Context) {
//     var user models.User
//     var userrequest keycloak.CreateUserRequest

//     // 1. Bind JSON input -> user
//     if err := c.ShouldBindJSON(&user); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
//         return
//     }

//     // 2. Save user ลง database ผ่าน service
//     if err := h.userService.RegisterUser(c.Request.Context(), &user); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user", "details": err.Error()})
//         return
//     }
//     if err := h.keycloakService.CreateUser (c.Request.Context(), &userrequest); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "not in keycloak", "details": err.Error()})
//         return
//     }
//     // 3. Return response
//     c.JSON(http.StatusCreated, gin.H{
//         "message": "User registered successfully",
//         "user_id": user.UserID, 
//     })
// }