package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"netforemost/pkg/token"
)

func TestAuthenticator(t *testing.T) {
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	parser := &token.JWT{
		AccessTokenSecretKey:        "SECRET",
		RefreshTokenSecretKey:       "SECRET2",
		Issuer:                      "Test",
		AccessTokenExpirationHours:  -1,
		RefreshTokenExpirationHours: -1,
	}
	const (
		role = Admin
		id   = "1"
	)
	tokenExpired, _, err := parser.Generate(role, id)
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	parser.AccessTokenExpirationHours = 1
	tkn, _, err := parser.Generate(role, id)
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	tests := []struct {
		name       string
		statusCode int
		token      string
	}{
		{
			name:       "Success",
			statusCode: http.StatusOK,
			token:      "Bearer " + tkn,
		},
		{
			name:       "Failure/InvalidFormat",
			statusCode: http.StatusUnauthorized,
			token:      tkn,
		},
		{
			name:       "Failure/InvalidToken",
			statusCode: http.StatusUnauthorized,
			token:      "Bearer x" + tkn,
		}, {
			name:       "Failure/TokenExpired",
			statusCode: http.StatusUnauthorized,
			token:      "Bearer " + tokenExpired,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}

			r.Header.Set("Authorization", test.token)

			h := Authenticator(parser)(f)

			mux := chi.NewMux()
			mux.Handle("/", h)
			mux.ServeHTTP(w, r)
			fmt.Println(w.Result())

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
		})
	}
}

func TestRefreshAuthenticator(t *testing.T) {
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	parser := &token.JWT{
		RefreshTokenSecretKey:       "SECRET2",
		Issuer:                      "Test",
		RefreshTokenExpirationHours: -1,
	}
	const (
		role = "1"
		id   = "1"
	)
	_, rTokenExpired, err := parser.Generate(role, id)
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	parser.RefreshTokenExpirationHours = 1
	_, rtoken, err := parser.Generate(role, id)
	if err != nil {
		assert.NoError(t, err)
		t.Fail()
	}

	tests := []struct {
		name       string
		statusCode int
		rtoken     string
	}{
		{
			name:       "Success",
			statusCode: http.StatusOK,
			rtoken:     "Bearer " + rtoken,
		},
		{
			name:       "Failure/InvalidFormat",
			statusCode: http.StatusUnauthorized,
			rtoken:     rtoken,
		},
		{
			name:       "Failure/InvalidToken",
			statusCode: http.StatusUnauthorized,
			rtoken:     "Bearer x" + rtoken,
		}, {
			name:       "Failure/TokenExpired",
			statusCode: http.StatusUnauthorized,
			rtoken:     "Bearer " + rTokenExpired,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				assert.NoError(t, err)
				t.Fail()
			}

			r.Header.Set("Authorization", test.rtoken)

			h := RefreshAuthenticator(parser)(f)

			mux := chi.NewMux()
			mux.Handle("/", h)
			mux.ServeHTTP(w, r)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
		})
	}
}
