package services

import (
	"github.com/anfelo/bookstore_users-api/domain/users"
	"github.com/anfelo/bookstore_users-api/utils/errors"
)

// CreateUser method incharge of creating a new user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
