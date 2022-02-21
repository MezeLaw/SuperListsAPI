package middleware

import (
	"SuperListsAPI/cmd/auth/models"
	"SuperListsAPI/cmd/auth/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ValidateJWTOnRequest(c *gin.Context) {

	//parsedToken := c.Request.Header["token"]

	parsedToken := c.Request.Header.Get("token")
	if parsedToken == "" {
		c.JSON(http.StatusUnauthorized, "missing token on request's header")
		c.Abort()
		return
	}

	//token := strings.Join(parsedToken, "")

	jwtWrapper := models.JwtWrapper{
		SecretKey:       repository.SECRET_KEY,
		Issuer:          repository.ISSUER,
		ExpirationHours: repository.EXPIRATION_HOURS,
	}

	claims, err := jwtWrapper.ValidateToken(parsedToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token present on request's header")
		c.Abort()
		return
	}

	userID := strconv.Itoa(int(claims.UserID))

	c.Request.Header.Add("role", claims.Role)
	c.Request.Header.Add("email", claims.Email)
	c.Request.Header.Add("user_id", userID)
	return
}
