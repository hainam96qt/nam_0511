package attendance

import (
	"context"
	"nam_0511/internal/model"
	"nam_0511/internal/repo/contracts"
	configs "nam_0511/pkg/config"
	"time"
)

type (
	Service struct {
		config configs.Config

		attendanceRepo AttendanceRepository
	}

	AttendanceRepository interface {
		CheckIn(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error
		GetListAttendanceByUserID(ctx context.Context, userID int64) ([]contracts.AttendanceContractAttendanceRecord, error)
	}
)

func NewAttendanceService(attendanceRepo AttendanceRepository) *Service {
	return &Service{
		attendanceRepo: attendanceRepo,
	}
}

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
		result = append(result, model.Attendance{
			UserID:      int(v.EmployeeId.Int64()),
			CheckInTime: time.Unix(v.CheckInTime.Int64(), 0),
			Location:    v.Details,
		})
	}
	return result
}
