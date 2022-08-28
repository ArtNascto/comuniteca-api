package middlewares

import (
	"net/http"

	service "comuniteca.com.br/api/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtService service.JWTService = service.JWTAuthService()
		token := ""

		tokenHeaders := c.Request.Header["Token"]
		if len(tokenHeaders) > 0 {
			token = tokenHeaders[0]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token não encontrado"})
			return
		}
		_, err := jwtService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		userName, err := jwtService.GetUser(token, c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		c.Set("user", userName)
	}
}
