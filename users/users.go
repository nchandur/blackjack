package users

import (
	"fmt"

	"github.com/nchandur/blackjack/model"
	"golang.org/x/crypto/bcrypt"
)

const BUY_IN int64 = 1000

type User struct {
	Username    string `json:"username"`
	Password    string
	SignedIn    bool `json:"signedIn"`
	model.Stats `json:"stats"`
	Kaasu       int64 `json:"kaasu"`
}

func NewUser(username, password string) (User, error) {
	if username == "" {
		return User{}, fmt.Errorf("failed to create user: username cannot be empty")
	}

	if password == "" {
		return User{}, fmt.Errorf("failed to create user: password cannot be empty")
	}

	user := User{
		Username: username,
		SignedIn: false,
		Kaasu:    BUY_IN,
	}

	hash, err := user.encryptPassword(password)

	if err != nil {
		return User{}, fmt.Errorf("failed to create player: %v", err)
	}

	user.Password = hash

	return user, nil

}

func (u User) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)

	if err != nil {
		return "", err
	}

	return string(hash), nil

}

func (u User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u User) String() string {

	return fmt.Sprintf("Username: %s\nSign In Status: %t\n%s\n\nKaasu: %d", u.Username, u.SignedIn, u.Stats.String(), u.Kaasu)

}

type UserManager interface {
	SignIn(username, password string) error
	SignOut(username, password string) error
	Create(username, password string) error
	Retrieve(username string) (User, error)
	Delete(username, password string) error
	Reset(username, password string) error
}
