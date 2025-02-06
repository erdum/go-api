package requests

type SignOnRequest struct {
	IdToken string `json:"idToken" validate:"required"`
}
