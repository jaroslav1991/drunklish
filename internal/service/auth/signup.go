package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/validator"
	"fmt"
)

func (a *Auth) SignUp(req dto.SignUpRequest) (*model.User, error) {
	if errDomain := validator.ValidateDomain(req.Email); errDomain != true {
		return nil, fmt.Errorf("invalid domain fail: %w", httputils.ErrValidation)
	}

	if errCountSymbol := validator.ValidateSymbol(req.Email); errCountSymbol != true {
		return nil, fmt.Errorf("invalid symbols: %w", httputils.ErrValidation)
	}

	if errLengthPassword := validator.LengthPassword(req.Password); errLengthPassword != true {
		return nil, fmt.Errorf("invalid length password: %w", httputils.ErrValidation)
	}

	existEmail, err := a.repo.ExistEmail(req.Email)
	if existEmail {
		return nil, httputils.ErrExistEmail
	}
	if err != nil {
		return nil, fmt.Errorf("check exist email: %w", httputils.ErrInternalServer)
	}

	hashPassword, err := a.hashFn(req.Password)
	if err != nil {
		return nil, fmt.Errorf("fail hash password: %w", httputils.ErrInternalServer)
	}

	createUser, err := a.repo.CreateUser(dto.SignUpRequest{
		Email:    req.Email,
		Password: hashPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return createUser, nil
}
