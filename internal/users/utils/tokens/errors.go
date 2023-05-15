package tokens

type ExpireTokenError struct {
	I18nId string
}

func NewExpireTokenError() *ExpireTokenError {
	return &ExpireTokenError{I18nId: "authorization.invalid_bearer_token"}
}

func (ExpireTokenError) Error() string {
	return "Token has expired"
}
