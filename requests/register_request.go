package requests

type RegisterRequest struct {
	Name			string `json:"name" validate:"required"`
	Email			string `json:"email" validate:"required,email"`
	PhoneNumber		string `json:"phone_number" validate:"required,max=15"`
	Password		string `json:"password" validate:"required,min=8"`
}
