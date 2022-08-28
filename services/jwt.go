package service

import (
	"fmt"
	"time"

	"comuniteca.com.br/api/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

//jwt service
type JWTService interface {
	GenerateToken(userName string) (string, int64)
	ValidateToken(token string) (*jwt.Token, error)
	GetUser(encodedToken string, c *gin.Context) (string, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := global.Config.JWT.Secret
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(userName string) (string, int64) {
	exp := time.Now().Add(time.Hour * 48).Unix()
	claims := &authCustomClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    userName,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t, exp
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}
func (service *jwtServices) GetUser(encodedToken string, c *gin.Context) (string, error) {
	var claims authCustomClaims

	_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(t *jwt.Token) (interface{}, error) {

		return []byte(service.secretKey), nil
	}, request.WithClaims(&claims))
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}
