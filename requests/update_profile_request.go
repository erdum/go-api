package requests

import "go-api/models"

type UpdateProfileRequest struct {
	Name			string `json:"name"`
	PhoneNumber		string `json:"phone_number"`
	Gender			models.Gender `json:"gender"`
	Avatar			string `json:"avatar"`
	Address			string `json:"address"`
	City			string `json:"city"`
	State			string `json:"state"`
	Country			string `json:"country"`
	ZipCode			string `json:"zip_code"`
	Lat				float32 `json:"lat"`
	Long			float32 `json:"long"`
}
