package services

import (
	"github.com/anfelo/bookstore_users-api/domain/users"
	"github.com/anfelo/bookstore_users-api/utils/errors"
)

// GetUser method incharge of retrieving user by id
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser method incharge of creating a new user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
