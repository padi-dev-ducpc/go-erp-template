package auth

import (
	"errors"
	"fmt"
	"time"

	db "github.com/deanrtaylor1/go-erp-template/db/sqlc"
	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JWTAuthenticator struct {
	secretKey string
}

func NewJWTAuthenticator(secretKey string) (Authenticator, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d chracters", minSecretKeySize)
	}
	return &JWTAuthenticator{secretKey: secretKey}, nil
}

func (a *JWTAuthenticator) CreateToken(email string, role db.UserRole, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, role, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(a.secretKey))
}

func (a *JWTAuthenticator) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(a.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
