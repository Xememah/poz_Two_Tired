package store

import (
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/converter"
	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
)

type Local struct {
	Data []model.Activity
}

func (l *Local) Init() error {
	data, err := converter.Load("./_data")
	if err != nil {
		return err
	}
	l.Data = data
	return nil
}

func narrowDataByTimestamp(timestamp time.Time, duration time.Duration, data []*model.Activity) []*model.Activity {
	for i := len(data) - 1; i >= 0; i-- {
		activity := data[i]

		matching := []model.Step{}
		for j, step := range activity.Steps {

			diff := timestamp.Unix() - step.Timestamp
			if diff >= 0 && float64(diff) <= duration.Seconds() {
				// always include a previous position when fetching timestamps
				if j > 0 && len(matching) == 0 {
					matching = append(matching, activity.Steps[j-1])
				}
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
