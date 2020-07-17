package models

import "time"

type WeightHistory struct {
	ID   string
	Date time.Time
	Max  int
	Min  int
}
