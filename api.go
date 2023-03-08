package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user in our system
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// createUser creates a new user
func createUser(name string, email string, password string) (*User, error) {
	// TODO: validate input data and create a new user object
	newUser := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	// TODO: save the user to a database or other storage medium
	return newUser, nil
}

// createUserHandler handles requests to create a new user
func createUserHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := createUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/users", createUserHandler)

	r.Run(":8000") // Start the server
}
