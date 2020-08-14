package services

import (
	"github.com/anfelo/bookstore_users-api/domain/users"
	"github.com/anfelo/bookstore_users-api/utils/crypto_utils"
	"github.com/anfelo/bookstore_users-api/utils/dates"
	"github.com/anfelo/bookstore_users-api/utils/errors"
)

var (
	// UsersService namespace variable
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

// GetUser method incharge of retrieving user by id
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser method incharge of creating a new user
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	user.CreatedAt = dates.GetNowStringDB()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser method incharge of updating a user
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	current, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser method incharge of deleting a user
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search method incharge of finding users that match a status
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

// LoginUser method incharge of login in a user
func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, *errors.RestErr) {
	user := &users.User{
		Email:    req.Email,
		Password: crypto_utils.GetMd5(req.Password),
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
