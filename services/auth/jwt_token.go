package auth

import (
	"go-api/config"
	"errors"
	"time"
	"reflect"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Token Token
	jwt.RegisteredClaims
}

type JWTToken struct {
	secretKey string
}

func NewJWTToken() *JWTToken {
	return &JWTToken{secretKey: config.GetConfig().Secret}
}

func (j *JWTToken) GenerateToken(tokenStruct Token) (string, error) {
	claims := TokenClaims{
		Token: tokenStruct,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTToken) ValidateToken(tokenString string) (
	*Token,
	error,
) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		fmt.Println(reflect.TypeOf(claims))
		// return claims.Token, nil
		return nil, nil
	}

	return nil, errors.New("Invalid token")
}
