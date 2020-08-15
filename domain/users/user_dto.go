package users

import (
	"strings"

	"github.com/anfelo/bookstore_utils/errors"
)

const (
	// StatusActive Active user status
	StatusActive = "active"
)

// User main user type
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

// Users list of users type
type Users []User

// Validate method that validates a user
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	return nil
}
