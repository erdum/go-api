package requests

type UpdatePasswordRequest struct {
	Email			string `json:"email" validate:"required,email"`
	NewPassword		string `json:"newPassword" validate:"required"`
}
