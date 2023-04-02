//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/token"
)

type Auth struct {
	repo            Repository
	hashFn          func(password string) (string, error)
	checkPasswordFn func(hash, password string) error
	generateTokenFn func(userId int64, email string) (string, error)
}

func NewAuthService(repo Repository) *Auth {
	return &Auth{repo: repo, hashFn: token.HashPassword}
}

type Repository interface {
	CreateUser(userDTO dto.SignUpRequest) (*model.User, error)
	CheckUserDB(user model.User) (*dto.ResponseUser, error)
	ExistEmail(email string) (bool, error)
}
