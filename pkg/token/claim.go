package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenNotGenerated = errors.New("token not generated")
	ErrParseToken        = errors.New("token can not parse")
	ErrSignatureInvalid  = errors.New("signature is invalid")
	ErrTokenExpired      = errors.New("token is expired")
)

// Claim is the user's claim on the token payload.
type Claim struct {
	*jwt.StandardClaims
	Role string `json:"role,omitempty"`
}

// JWT wraps the signing key and the issuer.
type JWT struct {
	AccessTokenSecretKey        string
	RefreshTokenSecretKey       string
	Issuer                      string
	AccessTokenExpirationHours  int64
	RefreshTokenExpirationHours int64
}

// Generate generate auth token.
func (j *JWT) Generate(role, id string) (string, string, error) {
	signingKey := []byte(j.AccessTokenSecretKey)
	now := time.Now()
	c := &Claim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(j.AccessTokenExpirationHours) * time.Hour).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.Issuer,
			Subject:   id,
			Audience:  "core",
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	accessToken, err := token.SignedString(signingKey)

	if err != nil {
		return "", "", err
	}

	signingKey = []byte(j.RefreshTokenSecretKey)
	c = &Claim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(j.RefreshTokenExpirationHours) * time.Hour).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.Issuer,
			Subject:   id,
			Audience:  "refreshToken",
		},
		Role: role,
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	refreshToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Parse parse token and validate.
func (j *JWT) Parse(tokenString string, refreshToken bool) (*Claim, error) {
	var signingKey []byte
	if refreshToken {
		signingKey = []byte(j.RefreshTokenSecretKey)
	} else {
		signingKey = []byte(j.AccessTokenSecretKey)
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, ErrSignatureInvalid
	} else if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		return nil, ErrParseToken
	}

	return claims, err
}
