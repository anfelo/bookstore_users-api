package users

import (
	"net/http"
	"strconv"

	"github.com/anfelo/bookstore_users-api/services"
	"github.com/anfelo/bookstore_users-api/utils/errors"

	"github.com/anfelo/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

// GetUser gets a user by user_id
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
	}
	user, getErr := services.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
