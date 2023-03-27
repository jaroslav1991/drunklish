package storage

import "time"

type (
	User struct {
		Id           int64  `json:"id" db:"id"`
		Email        string `json:"email" db:"email"`
		HashPassword string `json:"hash_password" db:"hash_password"`
	}

	Word struct {
		Word      string    `json:"word" db:"word"`
		Translate string    `json:"translate" db:"translate"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UserId    int64     `json:"user_id" db:"user_id"`
	}
)
