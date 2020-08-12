package users

import "encoding/json"

// PublicUser public user type
type PublicUser struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// PrivateUser private user type
type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// Marshal converts a list of users into a Public or Private list of users
func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

// Marshal converts a User into a Public or Private user
func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:        user.ID,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		}
	}
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
