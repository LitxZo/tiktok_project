package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtUserClaim struct {
	ID int
	jwt.RegisteredClaims
}

var signedKey = []byte(viper.GetString("Token.signedKey"))

func GenerateToken(id int) (string, error) {
	claim := JwtUserClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("Token.expiresTime") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(signedKey)
}

func ParseToken(tokenStr string) (JwtUserClaim, error) {
	claim := JwtUserClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil && token.Valid {
		err = errors.New("invalid token")
	}
	return claim, err
}

func TokenIsValid(tokenStr string) bool {
	claim := JwtUserClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil && token.Valid {
		return false
	}
	return true
}
