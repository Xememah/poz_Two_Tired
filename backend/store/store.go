package store

import (
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
)

type Query struct {
	// position from which to measure distance
	Position *model.Position
	// distance in metres
	Distance float64

	// relative timestamp
	Timestamp *time.Time
	// how long in to the future shall we look
	Offset time.Duration
}
type Store interface {
	Init() error
	Narrowed(query Query) ([]*model.Activity, error)
}
