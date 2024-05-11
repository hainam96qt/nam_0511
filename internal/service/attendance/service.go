package attendance

import (
	"context"
	"database/sql"
	db "nam_0801/internal/repo/dbmodel"
)

type (
	Service struct {
		userRepo UserRepository
	}

	AttendanceRepository interface {
		CheckIn(ctx context.Context, user db.CreateUserParams) (db.User, error)
	}
)

func NewUserService(ethereumConn *sql.DB, userRepo UserRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		userRepo:     userRepo,
	}
}
