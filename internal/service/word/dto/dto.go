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

type ResponseFromCreateWor struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type ResponseFromDeleting struct {
	Answer string `json:"answer"`
}

type RequestForGetByPeriod struct {
	UserId    int64  `json:"user_id"`
	CreatedAt Period `json:"created_at"`
}

type Period struct {
	FirstDate  time.Time `json:"first_date"`
	SecondDate time.Time `json:"second_date"`
}

type RequestForGettingWord struct {
	UserId int64 `json:"user_id"`
}

type RequestForDeletingWord struct {
	Word   string `json:"word"`
	UserId int64  `json:"user_id"`
}
