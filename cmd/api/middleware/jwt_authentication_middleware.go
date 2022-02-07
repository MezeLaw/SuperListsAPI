package middleware

import (
	"SuperListsAPI/cmd/auth/models"
	"SuperListsAPI/cmd/auth/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateJWTOnRequest(c *gin.Context) {

	parsedToken := c.Request.Header["token"]

	if parsedToken == nil {
		c.JSON(http.StatusUnauthorized, "missing token on request's header")
		c.Abort()
	}

	token := strings.Join(parsedToken, "")

	jwtWrapper := models.JwtWrapper{
		SecretKey:       repository.SECRET_KEY,
		Issuer:          repository.ISSUER,
		ExpirationHours: repository.EXPIRATION_HOURS,
	}

	claims, err := jwtWrapper.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token present on request's header")
		c.Abort()
	}

	c.Request.Header.Add("ROLE", claims.Role)
	c.Request.Header.Add("EMAIL", claims.Email)
	return
}
