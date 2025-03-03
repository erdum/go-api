package requests

type UpdatePasswordRequest struct {
	Email			string `json:"email" validate:"required,email"`
	NewPassword		string `json:"new_password" validate:"required"`
}
