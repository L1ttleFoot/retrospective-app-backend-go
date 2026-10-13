package core_service_tokens

import (
	"context"
	"time"
)

type ctxKeyClaims string

const claimsKey ctxKeyClaims = "claims"

type TokenType string

const (
	TypeAccess  TokenType = "acc"
	TypeRefresh TokenType = "ref"
)

type TokenClaims struct {
	UserID string
}

type TokenManager interface {
	Generate(claims TokenClaims, tokenType TokenType, ttl time.Duration) (string, error)
	Parse(token string, tokenType TokenType) (*TokenClaims, error)
}

func WithClaims(ctx context.Context, claims *TokenClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func FromContext(ctx context.Context) *TokenClaims {
	claims, ok := ctx.Value(claimsKey).(*TokenClaims)
	if !ok {
		panic("no claims in context")
	}

	return claims
}
