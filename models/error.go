package models

import "fmt"

// User
type ErrUserAlreadyExist struct {
	Email string
}

func IsErrUserAlreadyExist(err error) bool {
	_, ok := err.(ErrUserAlreadyExist)
	return ok
}

func (err ErrUserAlreadyExist) Error() string {
	return fmt.Sprintf("User already exists [name: %s].", err.Email)
}
