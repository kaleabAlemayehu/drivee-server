package dto

import "github.com/google/uuid"

type AuthResponse struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	MiddleName     string    `json:"middle_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture string    `json:"profile_picture"`
	Token          string    `json:"token"`
}

type LoginInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}
type ResetRequestInput struct {
	Email string `json:"email"`
}

type VerifyTokenInput struct {
	Token string `json:"token"`
}

type PasswordResetInput struct {
	Token    string    `json:"token"`
	UserID   uuid.UUID `json:"userID"`
	Password string    `json:"password"`
}
