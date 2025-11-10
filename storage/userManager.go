package storage

import (
	"fmt"

	"github.com/nchandur/blackjack/users"
)

type UserManager struct {
	users *[]users.User
}

func NewUserManager(users *[]users.User) UserManager {
	return UserManager{users: users}
}

func (u UserManager) Retrieve(username string) (users.User, error) {

	for _, play := range *(u.users) {
		if play.Username == username {
			return play, nil
		}
	}

	return users.User{}, fmt.Errorf("failed to retrieve users: users not found")

}

func (u UserManager) Create(username, password string) error {
	_, err := u.Retrieve(username)

	if err == nil {
		return fmt.Errorf("failed to create users: username %s already exists", username)
	}

	play, err := users.NewUser(username, password)

	if err != nil {
		return fmt.Errorf("failed to create users: %v", err)
	}

	*(u.users) = append(*(u.users), play)

	return nil

}

func (u UserManager) Delete(username, password string) error {
	idx := -1

	for i, play := range *(u.users) {
		if play.Username == username {
			if play.VerifyPassword(password) {
				idx = i
				break
			} else {
				return fmt.Errorf("failed to delete users: invalid password")
			}
		}
	}

	if idx == -1 {
		return fmt.Errorf("failed to delete users: %s does not exist", username)
	}

	*(u.users) = append((*(u.users))[:idx], (*(u.users))[idx+1:]...)

	return nil
}

func (u UserManager) SignIn(username, password string) error {

	success := false

	for idx := range *(u.users) {
		if (*(u.users))[idx].Username == username && (*(u.users))[idx].VerifyPassword(password) {
			(*(u.users))[idx].SignedIn = true
			success = true
		} else {
			(*(u.users))[idx].SignedIn = false

		}
	}

	if !success {
		return fmt.Errorf("username or password is incorrect")
	}

	return nil

}

func (u UserManager) SignOut(username, password string) error {

	success := false

	for idx := range *(u.users) {
		if (*(u.users))[idx].Username == username && (*(u.users))[idx].VerifyPassword(password) {
			(*(u.users))[idx].SignedIn = false
			success = true
		}
	}

	if !success {
		return fmt.Errorf("username or password is incorrect")
	}

	return nil
}

func (u UserManager) Reset(username, password string) error {

	success := false

	for idx := range *(u.users) {

		if (*(u.users))[idx].Username == username && (*(u.users))[idx].VerifyPassword(password) && (*(u.users))[idx].Kaasu == 0 {
			(*(u.users))[idx].Kaasu = users.BUY_IN
			success = true
		}

	}

	if !success {
		return fmt.Errorf("username or password is incorrect")
	}

	return nil

}
