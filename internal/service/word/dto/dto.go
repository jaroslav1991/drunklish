package dto

import "time"

type CreateWordRequest struct {
	Word      string    `json:"word"`
	Translate string    `json:"translate"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int64     `json:"user_id"`
}

type ResponseWord struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type ResponseWords struct {
	Words []ResponseWord
}

type RequestForGettingWord struct {
	UserId int64 `json:"user_id"`
}

type RequestForDeletingWord struct {
	Word   string `json:"word"`
	UserId int64  `json:"user_id"`
}
