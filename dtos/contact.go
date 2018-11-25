package dtos

// Contact -
type Contact struct {
	Dto
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
