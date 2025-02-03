package requests

type LoginRequest struct {
	IdToken string `json:"idToken" form:"idToken" validate:"required"`
}
