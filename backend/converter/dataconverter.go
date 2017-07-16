package converter

import (
	"io/ioutil"
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"

	"encoding/json"
	"os"
	"path"
)

type Collection struct {
	Features []Activity
}

type Activity struct {
	Type     model.ActivityType `json:"type"`
	Geometry struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		model.ActivityProperties
		Timestamps []int64   `json:"timestamps"`
		Elevations []float64 `json:"elevations"`
	} `json:"properties"`
}

func convert(data *Activity) model.Activity {
	converted := model.Activity{
		ActivityProperties: data.Properties.ActivityProperties,
		//Type:               data.Type,
	}
	for i, coords := range data.Geometry.Coordinates {
		step := model.Step{
			Timestamp: data.Properties.Timestamps[i],
			Position: model.Position{
				Longitude: coords[0],
				Latitude:  coords[1],
			},
		}
		if i == len(data.Geometry.Coordinates)-1 {
			step.Last = true
		}
		converted.Steps = append(converted.Steps, step)
	}
	converted.HashVal = converted.Hash()
	return converted
}

func Load(directory string) ([]model.Activity, int64, int64, error) {
	converted := []model.Activity{}
	files, err := ioutil.ReadDir(directory)
	min := time.Now().Unix()
	max := int64(0)
	if err != nil {
		return converted, min, max, err
	}
	for _, desc := range files {
		file, err := os.Open(path.Join(directory, desc.Name()))
		if err != nil {
			return converted, min, max, err
		}
		decoder := json.NewDecoder(file)
		raw := &Collection{}
		if err := decoder.Decode(raw); err != nil {
			return converted, min, max, err
		}
		for _, feature := range raw.Features {
			for _, ts := range feature.Properties.Timestamps {
				if ts > max {
					max = ts
				}
				if ts < min {
					min = ts
				}
			}
			converted = append(converted, convert(&feature))
		}
	}
	return converted, min, max, nil
}
