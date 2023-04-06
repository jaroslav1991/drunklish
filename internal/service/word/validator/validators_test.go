package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
