package errors

import "github.com/abolfazlalz/goasali/pkg/errors"

type UserHttpError struct {
	*errors.HttpError
}

func NewUserError(httpError *errors.HttpError) *UserHttpError {
	return &UserHttpError{HttpError: httpError}
}
