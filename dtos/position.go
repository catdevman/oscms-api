package dtos

// Position -
type Position struct {
	Dto
	Number   int    `json:"number"`
	Occupant string `json:"occupant"`
}
