package auth

type AuthService interface {
	AuthenticateWithThirdParty(idToken string) (map[string]string, error)
}
