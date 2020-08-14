package mysql_utils

import (
	"strings"

	"github.com/anfelo/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	// ErrorNoRows mysql no rows in result error
	ErrorNoRows = "no rows in result set"
)

// ParseError handle msql errors
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternatServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternatServerError("error processing request")
}
