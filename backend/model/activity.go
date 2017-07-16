package model

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type ActivityType string

const (
	ActivityTypeLine = "LineString"
)

type Activity struct {
	ActivityProperties `json:",omitempty"`
	HashVal            string `json:"hash,omitempty" db:"hash,omitempty"`
	//Type               ActivityType `json:"type,omitempty"`
	Steps []Step `json:"steps,omitempty"`
}

func (a *Activity) Hash() string {
	hash := sha1.Sum([]byte(fmt.Sprintf("%#v", a)))
	return hex.EncodeToString(hash[:])
}
