package token

import (
	"testing"
)

func TestClaim_GenerateToken(t *testing.T) {
	j := &JWT{
		AccessTokenSecretKey:        "SECRET",
		RefreshTokenSecretKey:       "SECRET2",
		Issuer:                      "test",
		AccessTokenExpirationHours:  3,
		RefreshTokenExpirationHours: 24,
	}
	const (
		role = "1"
		id   = "1"
	)

	atoken, rtoken, err := j.Generate(role, id)
	if err != nil {
		t.Errorf("Generate() error = %v, wantErr %v", nil, err)
	}
	if atoken == "" {
		t.Errorf("Generate() error = %v, wantErr %v", atoken, "")
	}
	if rtoken == "" {
		t.Errorf("Generate() error = %v, wantErr %v", rtoken, "")
	}
}

func TestClaim_Parse(t *testing.T) {
	tokenInvalidSignature := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJUZXN0Iiwic3ViIjoiMSIsIlJvbGUiOiIxIn0.zfTYduWZFn82zKNRGmOQ_HAGhFx1oWly8z98dekJToo"
	tokenExpired := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTExNDkzNTgsImlhdCI6MTYxMTE0OTM1OCwiaXNzIjoiVGVzdCIsInN1YiI6IjEiLCJSb2xlIjoiMSJ9.W00UI2IsYqKpj5hQxQmreYb4gpxYrdnUFxTDAq9X-Ts"

	j := &JWT{
		AccessTokenSecretKey:        "SECRET",
		RefreshTokenSecretKey:       "SECRET2",
		Issuer:                      "test",
		AccessTokenExpirationHours:  3,
		RefreshTokenExpirationHours: 24,
	}
	const (
		role = "1"
		id   = "1"
	)

	aTokenValid, rTokenValid, err := j.Generate(role, id)
	if err != nil {
		t.Fail()
	}

	tests := []struct {
		name   string
		atoken string
		rtoken string
		err    error
	}{
		{
			name:   "Success",
			atoken: aTokenValid,
			rtoken: rTokenValid,
			err:    nil,
		},
		{
			name:   "Invalid-Token",
			atoken: rTokenValid,
			rtoken: aTokenValid,
			err:    ErrSignatureInvalid,
		},
		{
			name:   "Invalid-Signature",
			atoken: tokenInvalidSignature,
			rtoken: tokenInvalidSignature,
			err:    ErrSignatureInvalid,
		},
		{
			name:   "Invalid-Signature",
			atoken: tokenInvalidSignature,
			rtoken: tokenInvalidSignature,
			err:    ErrSignatureInvalid,
		},
		{
			name:   "Expired",
			atoken: tokenExpired,
			rtoken: tokenExpired,
			err:    ErrSignatureInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := j.Parse(test.atoken, false)
			if err != test.err {
				t.Errorf("Parse() error = %v, wantErr %v", test.err, err)
			}
		})
	}
}
