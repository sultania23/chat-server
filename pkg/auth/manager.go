package auth

import "time"

type Token string

type TokenManager interface {
	GenerateToken(userId string, ttl time.Duration) (Token, error)
	ParseToken(accessToken Token) (string, error)
}
