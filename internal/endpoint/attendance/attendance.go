package attendance

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0511/internal/model"
	error2 "nam_0511/pkg/error"
	"nam_0511/pkg/midleware"
	"nam_0511/pkg/util/request"
	"nam_0511/pkg/util/response"
	"net/http"
)

type (
	Endpoint struct {
		authSvc AttendanceService
	}

	AttendanceService interface {
		CheckIn(ctx context.Context, userID int64, req *model.CheckInRequest) error
		GetListAttendanceByUserID(ctx context.Context, userId int64) (*model.GetListAttendanceByUserIDResponse, error)
		//UpdateTime(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse2, error)

	}
)

func InitAttendanceHandler(r *chi.Mux, authSvc AttendanceService) {
	authEndpoint := &Endpoint{
		authSvc: authSvc,
	}
	r.Route("/api/attendance", func(r chi.Router) {
		r.Use(midleware.Auth.ValidateRoleUser)

		r.Post("/check-in", authEndpoint.checkIn)
		r.Get("/", authEndpoint.getListAttendance)

		//r.Post("/update-time", authEndpoint.updateTime)

	})
}

func (e *Endpoint) checkIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CheckInRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)

	err := e.authSvc.CheckIn(ctx, int64(userIDCtx), &req)
	if err != nil {
		log.Printf("failed to check in : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func (e *Endpoint) getListAttendance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.LoginRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)

	res, err := e.authSvc.GetListAttendanceByUserID(ctx, int64(userIDCtx))
	if err != nil {
		log.Printf("failed to get list attendance : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

//func (e *Endpoint) updateTime(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	var req model.LoginRequest
//	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
//		log.Printf("read request body error: %s \n", err)
//		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
//		return
//	}
//
//	res, err := e.authSvc.UpdateTime(ctx, &req)
//	if err != nil {
//		log.Printf("failed to update time : %s \n", err)
//		response.Error(w, err)
//		return
//	}
//
//	response.JSON(w, http.StatusCreated, res)
//}
