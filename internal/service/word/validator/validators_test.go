package validator

import (
	"drunklish/internal/service/word/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCheckLengthWordAndTranslate1(t *testing.T) {
	type args struct {
		word      string
		translate string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive",
			args: args{
				word:      "qwe",
				translate: "qwe",
			},
			want: true,
		},
		{
			name: "negativeWord",
			args: args{
				word:      "",
				translate: "qwe",
			},
			want: false,
		},
		{
			name: "negativeTranslate",
			args: args{
				word:      "qwe",
				translate: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckLengthWordAndTranslate(tt.args.word, tt.args.translate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestCheckPlacesFirstOrSecondDate_PositiveInversion(t *testing.T) {
	testTime, err := time.Parse(time.RFC3339, "2050-04-29T00:00:01.000Z")
	assert.NoError(t, err)
	now := time.Now()
	var period dto.RequestForGetByPeriod

	expectedTime := dto.RequestForGetByPeriod{
		Token: period.Token,
		CreatedAt: dto.Period{
			FirstDate:  now,
			SecondDate: testTime,
		},
	}

	actual := CheckPlacesFirstOrSecondDate(expectedTime)
	assert.Equal(t, expectedTime, actual)
}

func TestCheckPlacesFirstOrSecondDate_NegativeInversion(t *testing.T) {
	testTime, err := time.Parse(time.RFC3339, "2050-04-29T00:00:01.000Z")
	assert.NoError(t, err)
	now := time.Now()
	var period dto.RequestForGetByPeriod

	expectedTime := dto.RequestForGetByPeriod{
		Token: period.Token,
		CreatedAt: dto.Period{
			FirstDate:  testTime,
			SecondDate: now,
		},
	}

	actual := CheckPlacesFirstOrSecondDate(expectedTime)
	assert.NotEqual(t, expectedTime, actual)
}
