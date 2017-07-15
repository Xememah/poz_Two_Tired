package converter

import (
	"io/ioutil"

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

func convert(data Activity) model.Activity {
	converted := model.Activity{
		ActivityProperties: data.Properties.ActivityProperties,
		//Type:               data.Type,
	}
	for i, coords := range data.Geometry.Coordinates {
		converted.Steps = append(converted.Steps, model.Step{
			Timestamp: data.Properties.Timestamps[i],
			Position: model.Position{
				Longitude: coords[1],
				Latitude:  coords[0],
			},
		})
	}
	converted.HashVal = converted.Hash()
	return converted
}

func Load(directory string) ([]model.Activity, error) {
	converted := []model.Activity{}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return converted, err
	}
	for _, desc := range files {
		file, err := os.Open(path.Join(directory, desc.Name()))
		if err != nil {
			return converted, err
		}
		decoder := json.NewDecoder(file)
		raw := &Collection{}
		if err := decoder.Decode(raw); err != nil {
			return converted, err
		}
		for _, feature := range raw.Features {
			converted = append(converted, convert(feature))
		}
	}
	return converted, nil
}
