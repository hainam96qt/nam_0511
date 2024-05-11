package attendance

import (
	"context"
	"math/big"
	"nam_0511/internal/model"
	"nam_0511/internal/repo/contracts/attendance"
	"time"
)

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
