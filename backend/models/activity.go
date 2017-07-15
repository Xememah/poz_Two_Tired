package models

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
	ActivityProperties
	Type  ActivityType
	Steps []Step
}

func (a *Activity) Hash() string {
	hash := sha1.Sum([]byte(fmt.Sprintf("%#v", a)))
	return hex.EncodeToString(hash[:])
}
