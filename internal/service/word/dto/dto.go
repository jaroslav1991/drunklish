package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

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
	Words ResponseWordList `json:"words"`
}

type ResponseWordList []ResponseWord

func (w ResponseWordList) Value() (driver.Value, error) {
	return json.Marshal(w)
}

func (w *ResponseWordList) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(data, &w)
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

type RequestForTraining struct {
	Words      ResponseWords `json:"words"`
	Answers    ResponseWords `json:"answers"`
	WordsTotal int64         `json:"words_total"`
	Token      string        `json:"token"`
}

type ResponseForTraining struct {
	Id    int64            `json:"id"`
	Words ResponseWordList `json:"words"`
}

type RequestTrainingInfo struct {
	TrainingId int64  `json:"training_id"`
	Token      string `json:"token"`
}

type ResponseTrainingInfo struct {
	Words   ResponseWordList `json:"words"`
	Answers ResponseWordList `json:"answers"`
}

type ResponseStatistic struct {
	CorrectAnswers ResponseWords `json:"correct_answers"`
	WrongAnswers   ResponseWords `json:"wrong_answers"`
}

type RequestStatistic struct {
	TrainingId int64  `json:"training_id"`
	Token      string `json:"token"`
}

type RequestCreateStatistic struct {
	//Id             int64         `json:"id"`
	TrainingId     int64         `json:"training_id"`
	CorrectAnswers ResponseWords `json:"correct_answers"`
	WrongAnswers   ResponseWords `json:"wrong_answers"`
	Token          string        `json:"token"`
}

type ResponseCreateStatistic struct {
	Id int64 `json:"id"`
}

type RequestFromUpload struct {
	Words ResponseWords `json:"words"`
	//Token string        `json:"token"`
}
