package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/users"
	"drunklish/internal/service/auth/validator"
	"errors"
	"fmt"
)

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id, email, hash_password`
)

var (
	ErrDomain         = errors.New("wrong domain for email")
	ErrSymbol         = errors.New("must be one '@' symbol in email")
	ErrLengthPassword = errors.New("password must be more than 5 symbols")
	ErrExistEmail     = errors.New("email is already exists")
)

func (a *Auth) SignUp(user *model.User) (*model.User, error) {
	if errDomain := validator.ValidateDomain(user.Email); errDomain != true {
		return nil, fmt.Errorf("%w", ErrDomain)
	}

	if errCountSymbol := validator.ValidateSymbol(user.Email); errCountSymbol != true {
		return nil, fmt.Errorf("%w", ErrSymbol)
	}

	if errLengthPassword := validator.LengthPassword(user.HashPassword); errLengthPassword != true {
		return nil, fmt.Errorf("%w", ErrLengthPassword)
	}

	if existEmail := validator.ExistEmail(a.db, user.Email); existEmail != true {
		return nil, fmt.Errorf("%w", ErrExistEmail)
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
