package model

type Position struct {
	Longitude float64 `json:"lng,omitempty" db:"longitude"`
	Latitude  float64 `json:"lat,omitempty" db:"latitude"`
}
