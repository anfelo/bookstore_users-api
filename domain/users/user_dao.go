package users

import (
	"fmt"

	"github.com/anfelo/bookstore_users-api/utils/dates"

	"github.com/anfelo/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get method retrieves a user from the DB by its id
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

// Save method saves a user to the DB
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(
				fmt.Sprintf("user %s already registered", user.Email),
			)
		}
		return errors.NewBadRequestError(
			fmt.Sprintf("user %d already exists", user.ID),
		)
	}

	user.CreatedAt = dates.GetNowString()
	usersDB[user.ID] = user
	return nil
}
