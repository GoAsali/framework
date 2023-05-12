package repository

import (
	"errors"
	"fmt"
)

var UserAlreadyExistsError = errors.New("not implement 'error' as some methods are missing: Error() string")

type HashPasswordError struct {
	err error
}

func (hpe HashPasswordError) Error() string {
	return fmt.Sprintf("error during hash password error: %v", hpe.err)
}
