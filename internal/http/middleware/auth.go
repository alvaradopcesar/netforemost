package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"netforemost/pkg/response"
	"netforemost/pkg/token"
)

type key string

const (
	IDKey   key = "id"
	RoleKey key = "role"
)

const (
	Admin   = "admin"
	Manager = "manager"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenNotFound = errors.New("token not found")
)

// Authenticator is an authentication middleware.
func Authenticator(parser token.Parser, validRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			tkn, err := tokenFromAuthorization(auth)
			if err != nil {
				_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			claim, err := parser.Parse(tkn, false)
			if err != nil {
				_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			if len(validRoles) > 0 && !roleAuth(claim.Role, validRoles...) {
				_ = response.HTTPError(w, r, http.StatusUnauthorized, "invalid role")
				return
			}

			ctx := context.WithValue(r.Context(), IDKey, claim.Subject)
			ctx = context.WithValue(ctx, RoleKey, claim.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RefreshAuthenticator is an authentication middleware for refresh token.
func RefreshAuthenticator(parser token.Parser) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			tkn, err := tokenFromAuthorization(auth)
			if err != nil {
				_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			claim, err := parser.Parse(tkn, true)
			if err != nil {
				_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), IDKey, claim.Subject)
			ctx = context.WithValue(ctx, RoleKey, claim.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Authenticator is an authentication middleware.
func roleAuth(roleName string, validRoles ...string) bool {
	roleName = strings.ToLower(roleName)
	for _, v := range validRoles {
		if v == roleName {
			return true
		}
	}

	return false
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", ErrTokenNotFound
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", ErrInvalidToken
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", ErrInvalidToken
	}

	return l[1], nil
}
