package store

import (
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
)

type Store interface {
	Init() error
	TimestampNarrowed(timestamp time.Time, offset time.Duration) ([]*model.Activity, error)
}
