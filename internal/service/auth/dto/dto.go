package dto

import "drunklish/internal/model"

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	User model.User
	//Token string
}

type ResponseFromSignIn struct {
	Token string
}
