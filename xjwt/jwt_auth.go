package xjwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

var defaultJwtSecretKey = "8ni3q2ruj092r4fj490&8^@gag"
var defaultJwtExp int64 = 24 * 60 * 60 //second

func SetSecret(secretKey string, exp int64) {
	defaultJwtSecretKey = secretKey
	defaultJwtExp = exp
}

type JWT struct {
	SigningKey []byte
	JwtExp     int64
}

type UserClaims struct {
	Id    uint   `json:"id"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func NewUserJWT() *JWT {
	return &JWT{
		SigningKey: []byte(defaultJwtSecretKey),
		JwtExp:     defaultJwtExp,
	}
}

func NewCustomJWT(exp int64) *JWT {
	return &JWT{
		SigningKey: []byte(defaultJwtSecretKey + "nufh9ur"),
		JwtExp:     exp,
	}
}

func (j *JWT) CreateToken(user *UserClaims) (string, error) {
	user.ExpiresAt = jwt.TimeFunc().Unix() + j.JwtExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		e := err.(*jwt.ValidationError)
		if e.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, TokenExpired
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func GetCurrentUser(c *gin.Context) *UserClaims {
	claims, ok := c.Get("claims")
	if !ok {
		return nil
	}
	user, ok := claims.(*UserClaims)
	if ok {
		return user
	}
	return nil
}
