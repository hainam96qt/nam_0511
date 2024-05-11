package attendance

import (
	"context"
	"math/big"
	"nam_0511/internal/model"
	contracts "nam_0511/internal/repo/contracts/attendance"
	"time"
)

func (s *Service) CheckOut(ctx context.Context, userID int64, req *model.CheckOutRequest) error {
	var attendance = contracts.AttendanceContractAttendanceRecord{
		EmployeeId:   big.NewInt(userID),
		Index:        big.NewInt(int64(req.Index)),
		CheckOutTime: big.NewInt(time.Now().Unix()),
		Details:      req.Location,
	}
	err := s.attendanceRepo.CheckOut(ctx, &attendance)
	if err != nil {
		return err
	}

	return nil
}
