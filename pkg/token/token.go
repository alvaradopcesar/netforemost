package token

// Token handles operations with JWT.
type Token interface {
	Parser
	Generator
}
