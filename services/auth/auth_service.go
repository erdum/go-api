package auth

import "go-api/requests"

type AuthService interface {
	Register(*requests.RegisterRequest) (map[string]string, error)
	Login(*requests.LoginRequest) (map[string]string, error)
	SignOn(*requests.SignOnRequest) (
		map[string]string,
		error,
	)
}
