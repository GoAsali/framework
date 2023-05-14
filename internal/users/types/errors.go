package types

import (
	pkgErrors "github.com/abolfazlalz/goasali/pkg/errors"
)

type NotValidBearerTokenError struct {
	*pkgErrors.I18nMessageError
}

func (NotValidBearerTokenError) Error() string {
	return "the token format entered is invalid"
}

type UserHttpError struct {
	*pkgErrors.HttpError
}

func NewUserError(httpError *pkgErrors.HttpError) *UserHttpError {
	return &UserHttpError{HttpError: httpError}
}
