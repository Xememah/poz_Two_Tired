package model

import (
	"math"
)

type Position struct {
	Longitude float64 `json:"lng,omitempty" db:"longitude"`
	Latitude  float64 `json:"lat,omitempty" db:"latitude"`
}

func (pos *Position) Distance(other Position) float64 {
	φ1 := toRadians(pos.Latitude)
	φ2 := toRadians(other.Latitude)
	Δφ := toRadians(other.Latitude - pos.Latitude)
	Δλ := toRadians(other.Longitude - pos.Longitude)

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return 6371000 * c
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
