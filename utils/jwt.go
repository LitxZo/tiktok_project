package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtUserClaim struct {
	ID   int
	time time.Time
	jwt.RegisteredClaims
}

var signedKey = []byte(viper.GetString("Token.signedKey"))

func GenerateToken(id int, t time.Time) (string, error) {
	expiresTime := viper.GetInt("Token.expiresTime")
	claim := JwtUserClaim{
		ID:   id,
		time: t,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(signedKey)
}

func ParseToken(tokenStr string) (JwtUserClaim, error) {
	claim := JwtUserClaim{}
	_, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("invalid token")
	}
	return claim, err
}

func TokenIsValid(tokenStr string) bool {
	claim := JwtUserClaim{}
	_, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}

func ParseTokenForId(tokenStr string) (int, error) {
	claim := JwtUserClaim{}
	_, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("invalid token")
	}
	return claim.ID, err
}
