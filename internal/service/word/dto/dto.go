package dto

import "time"

type CreateWordRequest struct {
	Word      string    `json:"word"`
	Translate string    `json:"translate"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}

type ResponseWord struct {
	Id        int64  `json:"id"`
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type ResponseWords struct {
	Words []ResponseWord `json:"words"`
}

type ResponseFromCreateWord struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type ResponseFromDeleting struct {
	Answer string `json:"answer"`
}

type RequestForGetByPeriod struct {
	Token     string `json:"token"`
	CreatedAt Period `json:"created_at"`
}

type Period struct {
	FirstDate  time.Time `json:"first_date"`
	SecondDate time.Time `json:"second_date"`
}

type RequestForGettingWord struct {
	Token string `json:"token"`
}

type RequestForDeletingWord struct {
	Token string `json:"token"`
	Id    int64  `json:"id"`
}

type RequestForUpdateWord struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
	Id        int64  `json:"id"`
	Token     string `json:"token"`
}
