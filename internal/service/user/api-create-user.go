package user

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	"nam_0801/pkg/util/password"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) error {
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}
	_, err = s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
