package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/users"
	"drunklish/internal/service/auth/validator"
	"errors"
	"fmt"
)

var (
	ErrDomain         = errors.New("wrong domain for email")
	ErrSymbol         = errors.New("must be one '@' symbol in email")
	ErrLengthPassword = errors.New("password must be more than 5 symbols")
	ErrExistEmail     = errors.New("email is already exists")
)

func (a *Auth) SignUp(req dto.SignUpRequest) (*model.User, error) {
	if errDomain := validator.ValidateDomain(req.Email); errDomain != true {
		return nil, fmt.Errorf("%w", ErrDomain)
	}

	if errCountSymbol := validator.ValidateSymbol(req.Email); errCountSymbol != true {
		return nil, fmt.Errorf("%w", ErrSymbol)
	}

	if errLengthPassword := validator.LengthPassword(req.Password); errLengthPassword != true {
		return nil, fmt.Errorf("%w", ErrLengthPassword)
	}

	if existEmail := validator.ExistEmail(a.db, req.Email); existEmail != true {
		return nil, fmt.Errorf("%w", ErrExistEmail)
	}

	hashPassword, err := users.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	createUser, err := a.repo.CreateUser(dto.SignUpRequest{
		Email:    req.Email,
		Password: hashPassword,
	})
	if err != nil {
		return nil, err
	}

	return createUser, nil
}
