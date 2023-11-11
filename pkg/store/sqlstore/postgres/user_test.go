package postgres

import (
	"testing"

	"github.com/kamogelosekhukhune777/tdd-api/pkg/domain"
)

func TestCreateUser(t *testing.T) {
	pStore := NewPostgresStore(nil)
	oldPassword := "password"

	user := &domain.User{
		Email:    "test@test.com",
		Password: "password",
		Name:     "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Errorf("want id not to be zero")
	}
	if user.Name != createdUser.Name {
		t.Errorf("expected %q: got %q", user.Name, createdUser.Name)
	}
	if createdUser.Password == oldPassword {
		t.Errorf("password not hashed")
	}
}
