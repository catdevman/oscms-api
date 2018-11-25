package dtos

// Grave -
type Grave struct {
	Dto
	ID        string     `json:"id"`
	Cemetery  string     `json:"cemetery"`
	Location  string     `json:"location"`
	Owners    []string   `json:"owners"`
	Positions []Position `json:"positions"`
}
