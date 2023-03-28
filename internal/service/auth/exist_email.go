package auth

import "drunklish/internal/service"

const (
	getEmailQuery = `select email from users where email=$1`
)

func (a *Auth) ExistEmail(db service.DB, email string) error {
	if err := db.QueryRowx(getEmailQuery, email).Scan(&email); err != nil {
		return err
	}

	return nil
}
