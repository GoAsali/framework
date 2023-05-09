package errors

import "github.com/nicksnyder/go-i18n/v2/i18n"

type HttpError struct {
	*i18n.Bundle
}
