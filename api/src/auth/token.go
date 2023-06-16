package auth

import (
	"api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userID uint64) (string, error) {
	roles := jwt.MapClaims{}
	roles["authorized"] = true
	roles["exp"] = time.Now().Add(time.Hour * 6).Unix()
	roles["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, roles)
	return token.SignedString([]byte(config.SecretKey))
}
