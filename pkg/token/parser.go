package token

// Parser parse a token string and returns the claim.
type Parser interface {
	Parse(tokenString string, refreshToken bool) (*Claim, error)
}
