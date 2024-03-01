package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTTokenManager struct {
	signingKey string
}

func NewJWTTokenManager(signingKey string) *JWTTokenManager {
	return &JWTTokenManager{
		signingKey: signingKey,
	}
}

func (m *JWTTokenManager) GenerateToken(userId string, ttl time.Duration) (Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})
	var str, err = token.SignedString([]byte(m.signingKey))
	return Token(str), err
}

func (m *JWTTokenManager) ParseToken(accessToken Token) (string, error) {
	token, err := jwt.Parse(string(accessToken), func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
