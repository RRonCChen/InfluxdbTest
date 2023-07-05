package model

import "time"

type Record struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]any
	Time        time.Time
}
