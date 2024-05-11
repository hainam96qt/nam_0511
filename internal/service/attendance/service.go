package attendance

import (
	"context"
	"nam_0511/internal/repo/contracts/attendance"
	configs "nam_0511/pkg/config"
)

type (
	Service struct {
		config configs.Config

		attendanceRepo AttendanceRepository
	}

	AttendanceRepository interface {
		CheckIn(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error
		CheckOut(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error
		GetListAttendanceByUserID(ctx context.Context, userID int64) ([]contracts.AttendanceContractAttendanceRecord, error)
		UpdateAttendanceTime(ctx context.Context, c contracts.AttendanceContractAttendanceRecord) error
	}
)

func NewAttendanceService(attendanceRepo AttendanceRepository) *Service {
	return &Service{
		attendanceRepo: attendanceRepo,
	}
}
