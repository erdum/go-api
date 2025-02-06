package auth

type Token struct {
	EntityID		uint
	EntityType		string
}

type TokenService interface {
	GenerateToken(token Token) (string, error)
	ValidateToken(tokenString string) (*Token, error)
}
