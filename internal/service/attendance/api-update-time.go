package attendance

import (
	"context"
	"math/big"
	"nam_0511/internal/model"
	contracts "nam_0511/internal/repo/contracts/attendance"
)

func (s *Service) UpdateCheckInTime(ctx context.Context, userID int64, req *model.UpdateTimeAttendanceRequest) error {
	var attendance = contracts.AttendanceContractAttendanceRecord{
		EmployeeId:  big.NewInt(userID),
		Index:       big.NewInt(int64(req.Index)),
		CheckInTime: big.NewInt(req.CheckInTime.Unix()),
		Details:     req.Reason,
	}
	err := s.attendanceRepo.UpdateAttendanceTime(ctx, attendance)
	if err != nil {
		return err
	}

	return nil
}
