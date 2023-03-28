package auth

import "drunklish/internal/service"

type Auth struct {
	db service.DB
}

func NewAuthService(db service.DB) *Auth {
	return &Auth{db: db}
}
