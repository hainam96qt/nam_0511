package attendance

import (
	"context"
	"nam_0511/internal/model"
	contracts "nam_0511/internal/repo/contracts/attendance"
	"time"
)

func (s *Service) GetListAttendanceByUserID(ctx context.Context, userID int64) (*model.GetListAttendanceByUserIDResponse, error) {
	data, err := s.attendanceRepo.GetListAttendanceByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.GetListAttendanceByUserIDResponse{Attendances: convertData(data)}, nil
}

func convertData(arr []contracts.AttendanceContractAttendanceRecord) []model.Attendance {
	var result []model.Attendance
	for _, v := range arr {
		var checkOutTimePtr *time.Time
		if !time.Unix(v.CheckOutTime.Int64(), 0).Equal(time.Unix(0, 0)) {
			checkInTime := time.Unix(v.CheckOutTime.Int64(), 0)
			checkOutTimePtr = &checkInTime
		}

		result = append(result, model.Attendance{
			UserID:       int(v.EmployeeId.Int64()),
			Index:        int(v.Index.Int64()),
			CheckInTime:  time.Unix(v.CheckInTime.Int64(), 0),
			CheckOutTime: checkOutTimePtr,
			Location:     v.Details,
		})
	}
	return result
}
