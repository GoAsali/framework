package tokens

import (
	"errors"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	errors2 "github.com/abolfazlalz/goasali/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
	"time"
)

var (
	JwtNotValid = errors.New("not valid tokens token")
)

type Token struct {
	jwtKey []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func New() *Token {
	key := []byte(os.Getenv("APP_KEY"))
	return &Token{key}
}

func (j *Token) GenerateJwtToken(user *models.User) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the Token claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In Token, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the Token string
	tokenString, err := token.SignedString(j.jwtKey)

	if err != nil {
		log.Printf("GenerateJwtToken: %v", err)
	}

	return tokenString, err
}

func (j *Token) ValidateJwtToken(token string) (*Claims, error) {
	claims := &Claims{}

	// Parse the Token string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.jwtKey, nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, JwtNotValid
	}
	result, ok := tkn.Claims.(Claims)
	if !ok {
		return nil, errors2.NewI18nError("invalid_bearer_token")
	}
	return &result, nil
}
