package converter

import (
	"io/ioutil"

	"encoding/json"
	"os"
	"path"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/models"
)

type Collection struct {
	Features []Activity
}

type Activity struct {
	Type     models.ActivityType `json:"type"`
	Geometry struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		models.ActivityProperties
		Timestamps []uint    `json:"timestamps"`
		Elevations []float64 `json:"elevations"`
	} `json:"properties"`
}

func convert(data Activity) models.Activity {
	converted := models.Activity{
		ActivityProperties: data.Properties.ActivityProperties,
		Type:               data.Type,
	}
	for i, coords := range data.Geometry.Coordinates {
		converted.Steps = append(converted.Steps, models.Step{
			Timestamp: data.Properties.Timestamps[i],
			Longitude: coords[1],
			Latitude:  coords[0],
		})
	}
	return converted
}

func Load(directory string) ([]models.Activity, error) {
	converted := []models.Activity{}
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
