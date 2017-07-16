package model

type ActivityProperties struct {
	Distance     int     `json:"distance,omitempty" db:"distance"`
	Duration     int     `json:"duration,omitempty" db:"duration"`
	EnergyBurned int     `json:"energy_burned,omitempty" db:"energy_burned"`
	Altitude     int     `json:"altitude,omitempty" db:"altitude"`
	Ascent       int     `json:"ascent,omitempty" db:"ascent"`
	MaxSpeed     float64 `json:"max_speed,omitempty" db:"max_speed"`
	ActivityType string  `json:"activity_type,omitempty" db:"activity_type"`
}
