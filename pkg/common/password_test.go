package common

import "testing"

func TestPasswordHash_Error(t *testing.T) {
	longPassword := make([]byte, 73)
	_, err := PasswordHash(string(longPassword))
	if err == nil {
		t.Errorf("expected an error; got nil")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashed, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckPassword(password, hashed)
	if err != nil {
		t.Errorf("password verification failed; got %q", err)
	}
}
