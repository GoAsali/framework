package services

import (
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"log"
	"time"
)

type AuthServiceLogs struct {
	AuthServiceI
}

func NewAuthServiceLogs(i AuthServiceI) AuthServiceI {
	return &AuthServiceLogs{i}
}

func (asl *AuthServiceLogs) Login(user *models.User, username string, password string) (*Token, error) {
	prefix := log.Prefix()
	log.SetPrefix("[AuthServiceLogs][Login] ")

	now := time.Now()

	re, err := asl.AuthServiceI.Login(user, username, password)

	defer func(token *Token, err error, start time.Time) {
		if err != nil {
			log.Printf("Method completed with error in %v: %v", time.Since(start), err)
		} else {
			log.Printf("Method completed in %v", time.Since(start))
		}
		log.SetPrefix(prefix)
	}(re, err, now)

	return re, err
}

func (asl *AuthServiceLogs) CreateAccount(user *models.User) (*Token, error) {
	prefix := log.Prefix()
	log.SetPrefix("[AuthServiceLogs][CreateAccount] ")

	now := time.Now()

	re, err := asl.AuthServiceI.CreateAccount(user)

	defer func(token *Token, err error, start time.Time) {
		if err != nil {
			log.Printf("Method completed with error in %v: %v", time.Since(start), err)
		} else {
			log.Printf("Method completed in %v", time.Since(start))
		}
		log.SetPrefix(prefix)
	}(re, err, now)

	return re, err
}
