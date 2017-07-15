package model

type ActivityProperties struct {
	Distance     int     `json:"distance,omitempty"`
	Duration     int     `json:"duration,omitempty"`
	EnergyBurned int     `json:"energy_burned,omitempty"`
	Altitude     int     `json:"altitude,omitempty"`
	Ascent       int     `json:"ascent,omitempty"`
	MaxSpeed     float64 `json:"max_speed,omitempty"`
	ActivityType string  `json:"activity_type,omitempty"`
}
