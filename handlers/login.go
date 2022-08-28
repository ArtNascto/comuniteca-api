package handlers

import (
	"net/http"

	"comuniteca.com.br/api/dtos"
	service "comuniteca.com.br/api/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginDto dtos.Login
	err := c.BindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var jwtService service.JWTService = service.JWTAuthService()

	token, expirationTime := jwtService.GenerateToken(loginDto.UserName)
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"access_token": token,
		"expires_in":   expirationTime,
	})
}
