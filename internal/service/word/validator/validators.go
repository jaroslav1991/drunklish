package validator

import (
	"drunklish/internal/service/word/dto"
	"time"
)

func CheckLengthWordAndTranslate(word, translate string) bool {
	if len(word) < 1 || len(translate) < 1 {
		return false
	}
	return true
}

func CheckPlacesFirstOrSecondDate(period dto.RequestForGetByPeriod) dto.RequestForGetByPeriod {
	durationFromFirst := time.Since(period.CreatedAt.FirstDate)
	durationFromSecond := time.Since(period.CreatedAt.SecondDate)

	if durationFromFirst > durationFromSecond {

		return dto.RequestForGetByPeriod{
			Token: period.Token,
			CreatedAt: dto.Period{
				FirstDate:  period.CreatedAt.FirstDate,
				SecondDate: period.CreatedAt.SecondDate,
			}}
	}

	period.CreatedAt.FirstDate, period.CreatedAt.SecondDate = period.CreatedAt.SecondDate, period.CreatedAt.FirstDate
	return dto.RequestForGetByPeriod{
		Token: period.Token,
		CreatedAt: dto.Period{
			FirstDate:  period.CreatedAt.FirstDate,
			SecondDate: period.CreatedAt.SecondDate,
		}}
}
