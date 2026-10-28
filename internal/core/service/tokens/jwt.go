package core_service_tokens

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	accessSecret  []byte
	refreshSecret []byte
}

type JwtClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

func NewTokenManager(config Config) *Tokens {
	return &Tokens{accessSecret: []byte(config.AccessSecret), refreshSecret: []byte(config.RefreshSecret)}
}

func (t *Tokens) Generate(claims TokenClaims, tokenType TokenType, ttl time.Duration) (string, error) {

	jwtClaims := JwtClaims{
		UserID: claims.UserID,
		Roles:  claims.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   string(tokenType),
		},
	}

	var secret []byte

	if tokenType == TypeAccess {
		secret = t.accessSecret
	} else {
		secret = t.refreshSecret
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("sign token error: %w", err)
	}

	return signedToken, nil
}

func (t *Tokens) Parse(tokenString string, tokenType TokenType) (*TokenClaims, error) {

	var claims JwtClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if tokenType == TypeAccess {
			return t.accessSecret, nil
		}
		return t.refreshSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &TokenClaims{UserID: claims.UserID, Roles: claims.Roles}, nil
}
