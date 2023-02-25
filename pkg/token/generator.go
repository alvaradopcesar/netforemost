package token

// Generator generates a new token.
type Generator interface {
	Generate(role, id string) (string, string, error)
}
