package middlewares

import (
	"fmt"
	"github.com/abolfazlalz/goasali/internal/users/types"
	"github.com/abolfazlalz/goasali/internal/users/utils/tokens"
	"github.com/abolfazlalz/goasali/pkg/errors"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func IsAuthMiddleware(c *gin.Context) {
	prefix := log.Prefix()

	defer func(prefix string) {
		log.SetPrefix(prefix)
	}(prefix)

	log.SetPrefix(fmt.Sprintf("%s - (IsAuthMiddleware): ", prefix))

	bearToken := c.GetHeader("Authorization")
	db := routes.NewContext(c).DB
	fmt.Println(db)
	tokenPart := strings.Split(bearToken, " ")
	if len(tokenPart) < 2 {
		//return types.NotValidBearerTokenError
		return
	}

	token := tokenPart[1]
	tokenSrv := tokens.New()
	jwtToken, err := tokenSrv.ValidateJwtToken(token)
	if err != nil {
		log.Println(err)
		if verr, ok := err.(types.NotValidBearerTokenError); ok {
			errors.NewByContext(c).I18nErrorMessage(c, verr.I18nId)
			return
		}
		if verr, ok := err.(*tokens.ExpireTokenError); ok {
			errors.NewByContext(c).I18nErrorMessage(c, verr.I18nId)
			return
		}
	}

	username := jwtToken.Username
	log.Println("Username", username)
	//return nil
}
