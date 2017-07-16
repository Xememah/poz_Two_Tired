package store

import (
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/converter"
	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
)

type Local struct {
	Data         []model.Activity
	minTimestamp int64
	maxTimestamp int64
}

func (l *Local) Init() error {
	data, min, max, err := converter.Load("./_data")
	if err != nil {
		return err
	}
	l.Data = data
	l.minTimestamp = min
	l.maxTimestamp = max
	return nil
}

func narrowDataByHash(hash string, data []*model.Activity) []*model.Activity {
	for _, activity := range data {
		if activity.HashVal == hash {
			return []*model.Activity{activity}
		}
	}
	return []*model.Activity{}
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

func narrowDataByLocation(pos model.Position, distance float64, data []*model.Activity) []*model.Activity {
	for i := len(data) - 1; i >= 0; i-- {
		activity := data[i]
		// only check starting and ending position
		firstStep := activity.Steps[0]
		lastStep := activity.Steps[len(activity.Steps)-1]
		if firstStep.Distance(pos) > distance && lastStep.Distance(pos) > distance {
			data = append(data[:i], data[i+1:]...)
		}
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

func (l *Local) Narrowed(query Query) ([]*model.Activity, error) {
	data := l.dataCopy()
	if query.Hash != "" {
		data = narrowDataByHash(query.Hash, data)
	}
	if query.Timestamp != nil {
		data = narrowDataByTimestamp(*query.Timestamp, query.Offset, data)
	}
	if query.Position != nil {
		data = narrowDataByLocation(*query.Position, query.Distance, data)
	}
	return data, nil
}

func (l *Local) MinMaxTimestamps() (int64, int64, error) {
	return l.minTimestamp, l.maxTimestamp, nil
}
