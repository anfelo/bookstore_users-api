package users

import (
	"fmt"

	"github.com/anfelo/bookstore_users-api/utils/mysql_utils"

	"github.com/anfelo/bookstore_users-api/datasources/mysql/users_db"

	"github.com/anfelo/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser        = "INSERT INTO users(first_name, last_name, email, password, status, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser           = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE id=?;"
	queryUpdateUser        = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser        = "DELETE FROM users WHERE id=?;"
	queryFindUsersByStatus = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE status=?;"
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
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt,
	); err != nil {
		return mysql_utils.ParseError(err)
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

	insertResult, saveErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
	)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
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

// Update method updates a user in DB
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternatServerError(err.Error())
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	return nil
}

// Delete method deletes a user from DB
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternatServerError(err.Error())
	}
	defer stmt.Close()

	if _, delErr := stmt.Exec(user.ID); delErr != nil {
		return mysql_utils.ParseError(delErr)
	}

	return nil
}

// FindByStatus method finds users that match status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors.NewInternatServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternatServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt,
		); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
