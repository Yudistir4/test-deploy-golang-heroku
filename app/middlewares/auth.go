package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJwt) GenerateToken(userId int) string {
	claims := JwtCustomClaims{
		ID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))
	return token
}
