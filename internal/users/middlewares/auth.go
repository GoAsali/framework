package middlewares

import (
	"fmt"
	"github.com/abolfazlalz/goasali/internal/users/db/models"
	"github.com/abolfazlalz/goasali/internal/users/db/repository"
	"github.com/abolfazlalz/goasali/internal/users/utils/tokens"
	"github.com/abolfazlalz/goasali/pkg/errors"
	routes "github.com/abolfazlalz/goasali/pkg/http/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func IsAuthMiddleware(c *gin.Context) {
	prefix := log.Prefix()

	defer func(prefix string) {
		log.SetPrefix(prefix)
	}(prefix)

	log.SetPrefix(fmt.Sprintf("%s - (IsAuthMiddleware): ", prefix))

	bearToken := c.GetHeader("Authorization")
	context := routes.NewContext(c)
	db := context.DB

	tokenPart := strings.Split(bearToken, " ")
	if len(tokenPart) < 2 {
		//return types.NotValidBearerTokenError
		return
	}

	token := tokenPart[1]
	tokenSrv := tokens.New()
	jwtToken, err := tokenSrv.ValidateJwtToken(token)
	httpErr := errors.NewByContext(c)
	if err != nil {
		log.Printf("Error in check middleware auth: %v", err)
		if verr, ok := err.(errors.I18nMessageError); ok {
			errors.NewByContext(c).I18nErrorMessage(c, verr.I18nId)
			httpErr.HandleHttp(c, httpErr.HttpCode(http.StatusUnauthorized), httpErr.I18nErrorMessageConfig(c, verr.I18nId))
			return
		}

		errors.NewByContext(c).I18nErrorMessage(c, "errors.internal_server")
		return
	}

	userRepo := repository.NewUserRepository(db, context.Cache)
	user := models.User{}
	if err := userRepo.FindByUsername(jwtToken.Username, &user); err != nil {
		httpErr.HandleGinError(err, c)
		return
	}

	c.Set("user", &user)

	c.Next()
}
