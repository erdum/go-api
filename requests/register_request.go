package requests

type RegisterRequest struct {
	Name			string `json:"name" validate:"required"`
	Email			string `json:"email" validate:"required,email"`
	PhoneNumber		string `json:"phoneNumber" validate:"required,lte=15"`
	Password		string `json:"password" validate:"required,gte=8"`
}
