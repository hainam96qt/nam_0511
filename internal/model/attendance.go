package model

import "time"

type CheckInRequest struct {
	Location string `json:"location"`
}

type CheckOutRequest struct {
	Index    int    `json:"index"`
	Location string `json:"location"`
}

type GetListAttendanceByUserIDRequet struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

type GetListAttendanceByUserIDResponse struct {
	Attendances []Attendance `json:"attendances"`
}

type Attendance struct {
	UserID       int        `json:"user_id"`
	Index        int        `json:"index"`
	CheckInTime  time.Time  `json:"check_in_time"`
	CheckOutTime *time.Time `json:"check_out_time,omitempty"`
	Location     string     `json:"location"`
}

type UpdateTimeAttendanceRequest struct {
	Index       int       `json:"index"`
	CheckInTime time.Time `json:"check_in_time"`
	Reason      string    `json:"reason"`
}
