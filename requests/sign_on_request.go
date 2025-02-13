package requests

type SignOnRequest struct {
	IdToken string `json:"id_token" validate:"required"`
}
