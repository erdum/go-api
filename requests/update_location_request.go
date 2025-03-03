package requests

type UpdateLocationRequest struct {
	Lat			float32 `json:"lat" validate:"required"`
	Long		float32 `json:"long" validate:"required"`
	Location	string `json:"location" validate:"required"`
	City		string `json:"city" validate:"required"`
	State		string `json:"state" validate:"required"`
}
