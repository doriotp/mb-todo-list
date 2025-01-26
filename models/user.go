package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Country    string `json:"country"`
	Occupation string `json:"occupation"`
	Phone      string `json:"phone"`
}

type RegisterResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Occupation string `json:"occupation"`
	Phone      string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}
