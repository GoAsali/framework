package types

import (
	"fmt"
	pkgErrors "github.com/abolfazlalz/goasali/pkg/errors"
)

//var RoleAlreadyExists = errors.New("role already exists")

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

type RoleAlreadyExists struct {
	*pkgErrors.I18nMessageError
	Name string
}

func NewRoleAlreadyExists(name string) *RoleAlreadyExists {
	return &RoleAlreadyExists{
		Name:             name,
		I18nMessageError: &pkgErrors.I18nMessageError{I18nId: ""},
	}
}

func (rae RoleAlreadyExists) Error() string {
	return fmt.Sprintf("%s name already exists as a role name.", rae.Name)
}
