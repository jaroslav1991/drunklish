package users

import "testing"

func TestHashPassword(t *testing.T) {
	password1, err := HashPassword("1234")
	if err != nil {
		t.Error("fail on generate hash password", err)
	}

	password2, err := HashPassword("1234")
	if err != nil {
		t.Error("fail on generate hash password", err)
	}

	t.Log(password1)
	t.Log(password2)

}

func TestCheckPasswordHash(t *testing.T) {
	hash, err := HashPassword("1234")
	if err != nil {
		t.Error("fail on generate hash password", err)
	}

	t.Log(hash)
	result := CheckPasswordHash("1234", hash)
	if !result {
		t.Error("fail with hash")
	}
}
