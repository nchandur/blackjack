package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const BUY_IN int = 1000

type Round struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
	Pushed int `json:"pushed"`
}

type User struct {
	Username string `json:"username"`
	Password string
	SignedIn bool `json:"signedIn"`
	Round    `json:"rounds"`
	Kaasu    int `json:"kaasu"`
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
		Round:    Round{},
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

	return fmt.Sprintf("Username: %s\nSign In Status: %t\n\nStats: \nRounds Played: %d\nRounds Won: %d\nRounds Lost: %d\nRounds Pushed: %d\n\nKaasu: %d", u.Username, u.SignedIn, u.Played, u.Won, u.Lost, u.Pushed, u.Kaasu)

}

type UserManager interface {
	SignIn(username, password string) error
	SignOut(username, password string) error
	Create(username, password string) error
	Retrieve(username string) (User, error)
	Delete(username, password string) error
}
