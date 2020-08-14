package users

// LoginRequest login request struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
