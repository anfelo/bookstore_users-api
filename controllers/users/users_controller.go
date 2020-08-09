package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser gets a user by user_id
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implemented")
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implemented")
}
