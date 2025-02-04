package auth

import (
	"go-api/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(entityID uint, entityType string) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
}

type TokenClaims struct {
	EntityID   uint
	EntityType string
	jwt.RegisteredClaims
}

type JWTService struct {
	secretKey string
}

func NewJWTService() *JWTService {
	return &JWTService{secretKey: config.GetConfig().Secret}
}

func (j *JWTService) GenerateToken(
	entityID uint,
	entityType string,
) (string, error) {
	claims := TokenClaims{
		EntityID:   entityID,
		EntityType: entityType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(tokenString string) (*TokenClaims, error) {
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
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}
