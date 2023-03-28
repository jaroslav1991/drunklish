package validator

import "testing"

func TestValidateDomain(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validEmail",
			args: args{email: "test@gmail.com"},
			want: true,
		},
		{
			name: "validEmail2",
			args: args{email: "test@yahoo.com"},
			want: true,
		},
		{
			name: "invalidDomain",
			args: args{email: "test@yandex.com"},
			want: false,
		},
		{
			name: "invalidEmail",
			args: args{email: "@gmail.com"},
			want: false,
		},
		{
			name: "invalidEmail2",
			args: args{email: "test@yahoo.ru"},
			want: false,
		},
		{
			name: "invalidEmail3",
			args: args{email: "@yandex.ru"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateDomain(tt.args.email); got != tt.want {
				t.Errorf("ValidateDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateSymbol(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validCountSymbol",
			args: args{email: "test@yandex.ru"},
			want: true,
		},
		{
			name: "zeroSymbol",
			args: args{email: "test.yandex.ru"},
			want: false,
		},
		{
			name: "manySymbols",
			args: args{email: "test@@gmail.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateSymbol(tt.args.email); got != tt.want {
				t.Errorf("ValidateSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positivePassword",
			args: args{password: "test password"},
			want: true,
		},
		{
			name: "tooSmallPassword",
			args: args{password: "test1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthPassword(tt.args.password); got != tt.want {
				t.Errorf("LengthPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
