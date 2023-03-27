package storage

import "drunklish/internal/users"

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id, email, hash_password`
)

func (s *Storage) SignUp(user *User) (*User, error) {
	hashPassword, err := users.HashPassword(user.HashPassword)
	if err != nil {
		return nil, err
	}
	user.HashPassword = hashPassword

	if err := s.DB.QueryRowx(createUserQuery, user.Email, user.HashPassword).Scan(&user.Id, &user.Email, &user.HashPassword); err != nil {
		return nil, err
	}
	return user, nil
}
