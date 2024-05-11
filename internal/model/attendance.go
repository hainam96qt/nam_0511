package model

import "time"

type CheckInRequest struct {
	Location string `json:"location"`
}

type GetListAttendanceByUserIDResponse struct {
	Attendances []Attendance `json:"attendances"`
}

type Attendance struct {
	UserID      int       `json:"user_id"`
	CheckInTime time.Time `json:"check_in_time"`
	Location    string    `json:"location"`
}
