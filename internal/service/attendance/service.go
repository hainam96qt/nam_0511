package attendance

import (
	"context"
	"math/big"
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
	}
)

func NewAttendanceService(attendanceRepo AttendanceRepository) *Service {
	return &Service{
		attendanceRepo: attendanceRepo,
	}
}

func (s *Service) CheckIn(ctx context.Context, userID int64, req *model.CheckInRequest) error {
	var attendance = contracts.AttendanceContractAttendanceRecord{
		EmployeeId:  big.NewInt(userID),
		CheckInTime: big.NewInt(time.Now().Unix()),
		Details:     req.Location,
	}
	err := s.attendanceRepo.CheckIn(ctx, &attendance)
	if err != nil {
		return err
	}

	return nil
}
