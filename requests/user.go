package requests

// UserRequest is a struct to hold the request body for creating a new user using http POST method
type RegisterUserRequest struct {
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}
