package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/users"
)

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id, email, hash_password`
)

func (a *Auth) SignUp(user *model.User) (*model.User, error) {
	hashPassword, err := users.HashPassword(user.HashPassword)
	if err != nil {
		return nil, err
	}
	user.HashPassword = hashPassword

	if err := a.db.QueryRowx(createUserQuery, user.Email, user.HashPassword).Scan(&user.Id, &user.Email, &user.HashPassword); err != nil {
		return nil, err
	}
	return user, nil
}
