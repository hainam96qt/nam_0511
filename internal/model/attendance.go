package model

import "time"

type CheckInRequest struct {
	Time     time.Time
	Location string
	Note     string
}
