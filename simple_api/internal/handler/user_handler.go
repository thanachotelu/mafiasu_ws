package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"users/internal/user"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users [get]
func GetAllUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := user.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags Users
// @Produce  json
// @Param   id   path      int     true  "User ID"
// @Success 200  {object}  User
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "name": "โอสธี สุขภูตานนท์"})
}

// DeleteUserHandler handles the delete user API
// @Summary      Delete user by ID
// @Description  Delete user from database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param   id   path      int     true  "User ID"
// @Success 200  {object}  map[string]string
// @Failure 404  {object}  ErrorResponse
// @Failure 400  {object}  ErrorResponse
// @Failure 500  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [delete]
func DeleteUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = user.DeleteUser(db, id)
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user, database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

// @Summary Update user information
// @Description Update details of a user
// @Tags Users
// @Produce  json
// @Param        id    path      int           true  "User ID"
// @Param        user  body      user.UserInput  true  "User Data"
// @Success 200  {object}  map[string]string
// @Failure 404  {object}  ErrorResponse
// @Router  /api/v1/users/{id} [put]
func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		var User user.UserInput

		if err := c.ShouldBindJSON(&User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input format",
			})
			return
		}

		updatedUser := user.User{
			ID:    id,
			Name:  User.Name,
			Email: User.Email,
		}

		err = user.UpdateUser(db, updatedUser)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "fail to updated ",
			})
			return
		}
		// Success response
		c.JSON(http.StatusCreated, gin.H{
			"message": "Updated successfully",
		})
	}
}

// AddUserHandler handles the add user API
// @Summary      Add a new user
// @Description  Adds a new user to the database
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      user.UserInput  true  "User Data"
// @Success      200   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /api/v1/users [post]
// AddUserHandler handles the add user API
func AddUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser user.User

		// Bind JSON input to the User struct
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input format",
			})
			return
		}

		// Call AddUser from user package
		err := user.AddUser(db, newUser)
		if err != nil {
			if err.Error() == "email already exists" {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Email already exists",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to add user",
				})
			}
			return
		}

		// Success response
		c.JSON(http.StatusCreated, gin.H{
			"message": "User added successfully",
		})
	}
}