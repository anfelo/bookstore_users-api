package users

import (
	"fmt"
	"strings"

	"github.com/anfelo/bookstore_users-api/utils/mysql_utils"

	"github.com/anfelo/bookstore_users-api/datasources/mysql/users_db"
	"github.com/anfelo/bookstore_users-api/logger"
	"github.com/anfelo/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, password, status, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, status, created_at FROM users WHERE email=? AND password=? AND status=?;"
)

// Get method retrieves a user from the DB by its id
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user stmt", err)
		return errors.NewInternatServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt,
	); err != nil {
		logger.Error("error when trying to scan result in get user", err)
		return errors.NewInternatServerError("database error")
	}

	return nil
}

// Save method saves a user to the DB
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user stmt", err)
		return errors.NewInternatServerError("database error")
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
		logger.Error("error when trying to exec save user", err)
		return errors.NewInternatServerError("database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last inserted id save user", err)
		return errors.NewInternatServerError("database error")
	}
	user.ID = userID
	return nil
}

// Update method updates a user in DB
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user stmt", err)
		return errors.NewInternatServerError("database error")
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if saveErr != nil {
		logger.Error("error when trying to exec update user", saveErr)
		return errors.NewInternatServerError("database error")
	}
	return nil
}

// Delete method deletes a user from DB
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare update user stmt", err)
		return errors.NewInternatServerError("database error")
	}
	defer stmt.Close()

	if _, delErr := stmt.Exec(user.ID); delErr != nil {
		logger.Error("error when trying to exec delete user", delErr)
		return errors.NewInternatServerError("database error")
	}

	return nil
}

// FindByStatus method finds users that match status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user stmt", err)
		return nil, errors.NewInternatServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to query user", err)
		return nil, errors.NewInternatServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt,
		); err != nil {
			logger.Error("error when trying to scan user row", err)
			return nil, errors.NewInternatServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

// FindByEmailAndPassword method retrieves a user from the DB by its id
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password stmt", err)
		return errors.NewInternatServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if err := result.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreatedAt,
	); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to scan result in find by email and password", err)
		return errors.NewInternatServerError("database error")
	}

	return nil
}
