package store

import (
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
)

type Local struct {
	Data []model.Activity
}

func narrowDataByTimestamp(timestamp time.Time, duration time.Duration, data []*model.Activity) []*model.Activity {
	for i := len(data) - 1; i >= 0; i-- {
		activity := data[i]
		matching := []model.Step{}
		for _, step := range activity.Steps {
			diff := timestamp.Unix() - step.Timestamp
			if diff >= 0 && float64(diff) <= duration.Seconds() {
				matching = append(matching, step)
			}
		}
		if len(matching) == 0 {
			data = append(data[:i], data[i+1:]...)
		}
		activity.Steps = matching
	}
	return data
}

func (l *Local) dataCopy() []*model.Activity {
	tmp := make([]*model.Activity, len(l.Data))
	for i, ac := range l.Data {
		act := &model.Activity{}
		*act = ac
		tmp[i] = act
	}
	return tmp
}

func (l *Local) TimestampNarrowed(timestamp time.Time, duration time.Duration) ([]*model.Activity, error) {
	return narrowDataByTimestamp(timestamp, duration, l.dataCopy()), nil
}
