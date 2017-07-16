package store

import (
	"fmt"
	"os"
	"time"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const stepsSchema = `CREATE TABLE twotired.steps (
		hash VARCHAR(40) NOT NULL ,
		timestamp INT(11) NOT NULL ,
		longitude DECIMAL(8,6) NOT NULL ,
		latitude DECIMAL(8,6) NOT NULL ,
		INDEX (hash), INDEX(timestamp))`
const activitiesSchema = `CREATE TABLE IF NOT EXIST twotired.activities (
		hash VARCHAR(40) NOT NULL ,
		distance INT NOT NULL ,
		duration INT NOT NULL ,
		altitude INT NOT NULL ,
		ascent INT NOT NULL ,
		max_speed DOUBLE NOT NULL ,
		energy_burned INT NOT NULL ,
		activity_type VARCHAR NOT NULL ,
		PRIMARY KEY (hash))`

type Postgres struct {
	DB *sqlx.DB
}

func (p *Postgres) Init() error {
	con, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@database/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		return fmt.Errorf("could not open database connection: %s", err.Error())
	}
	p.DB = con

	p.DB.MustExec(activitiesSchema)
	p.DB.MustExec(stepsSchema)
	return nil
}

func (p *Postgres) TimestampNarrowed(timestamp time.Time, offset time.Duration) ([]model.Activity, error) {
	return []model.Activity{}, nil
}
