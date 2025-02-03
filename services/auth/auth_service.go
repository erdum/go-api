package auth

type AuthService interface {
	AuthenticateWithThirdParty(idToken string) (interface{}, error)
}
