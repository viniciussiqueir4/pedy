package auth

import "pedy/models"

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
