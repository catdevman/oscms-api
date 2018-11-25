package dtos

// Cemetery -
type Cemetery struct {
	Dto
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
