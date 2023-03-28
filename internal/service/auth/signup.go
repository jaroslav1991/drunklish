package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/users"
	"drunklish/internal/service/auth/validator"
	"errors"
)

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id, email, hash_password`
)

func (a *Auth) SignUp(user *model.User) (*model.User, error) {
	if errDomain := validator.ValidateDomain(user.Email); errDomain != true {
		return nil, errors.New("wrong domain for email")
	}

	if errCountSymbol := validator.ValidateSymbol(user.Email); errCountSymbol != true {
		return nil, errors.New("must be one '@' symbol in email")
	}

	if errLengthPassword := validator.LengthPassword(user.HashPassword); errLengthPassword != true {
		return nil, errors.New("password must be more than 5 symbols")
	}

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
