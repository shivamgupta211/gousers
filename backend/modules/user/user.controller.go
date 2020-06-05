package usermodule

import (
	"errors"
)

// CreateUserController ... CreateUserController
func CreateUserController(data User) (User, error) {
	user, err := CreateUser(data)
	if err != nil {
		return User{}, errors.New("Error while creating User")
	}

	return user, nil
}
