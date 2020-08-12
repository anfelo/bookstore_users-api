package users

import (
	"fmt"
	"strings"

	"github.com/anfelo/bookstore_users-api/datasources/mysql/users_db"
	"github.com/anfelo/bookstore_users-api/utils/dates"

	"github.com/anfelo/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?;"
)

// Get method retrieves a user from the DB by its id
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternatServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt,
	); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.ID),
			)
		}
		return errors.NewInternatServerError(
			fmt.Sprintf("error when trying to retrieve user %d", user.ID),
		)
	}

	return nil
}

// Save method saves a user to the DB
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternatServerError(err.Error())
	}
	defer stmt.Close()

	user.CreatedAt = dates.GetNowString()
	insertResult, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.CreatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternatServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternatServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}
	user.ID = userID
	return nil
}
