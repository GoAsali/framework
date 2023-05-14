package errors

type I18nMessageError struct {
	error
	I18nId string
}

func NewI18nError(i18nId string) I18nMessageError {
	return I18nMessageError{
		I18nId: i18nId,
	}
}
