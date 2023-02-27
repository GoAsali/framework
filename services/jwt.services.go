package services

import (
	"asalpolaki/models"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTService struct{}

type JWTData struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (service JWTService) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashed)
}

func (service JWTService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service JWTService) GenerateJWT(user models.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println("secret key:", secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["user_username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 6).Unix()

	var mySigningKey = []byte(secretKey)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		_ = fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func (service JWTService) EncodeJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	var mySigningKey = []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Claims.(jwt.MapClaims); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	fmt.Println(err)
	return nil, err
}
