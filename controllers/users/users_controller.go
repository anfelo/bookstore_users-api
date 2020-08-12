package users

import (
	"net/http"
	"strconv"

	"github.com/anfelo/bookstore_users-api/services"
	"github.com/anfelo/bookstore_users-api/utils/errors"

	"github.com/anfelo/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

// Get gets a user by user_id
func Get(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, user.Marshal(c.GetHeader("X-Public") == "true"))
}

// Create creates a new user
func Create(c *gin.Context) {
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
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// Update updates a user
func Update(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	isPartial := c.Request.Method == http.MethodPatch
	user.ID = userID
	result, saveErr := services.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// Delete deletes a user
func Delete(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if delErr := services.DeleteUser(userID); delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search finds a user by status
func Search(c *gin.Context) {
	status := c.Query("status")

	result, getErr := services.Search(status)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}
