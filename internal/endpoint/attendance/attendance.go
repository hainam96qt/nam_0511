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
		attendanceSvc AttendanceService
	}

	AttendanceService interface {
		CheckIn(ctx context.Context, userID int64, req *model.CheckInRequest) error
		CheckOut(ctx context.Context, userID int64, req *model.CheckOutRequest) error
		GetListAttendanceByUserID(ctx context.Context, userId int64) (*model.GetListAttendanceByUserIDResponse, error)
		UpdateCheckInTime(ctx context.Context, userID int64, req *model.UpdateTimeAttendanceRequest) error
	}
)

func InitAttendanceHandler(r *chi.Mux, authSvc AttendanceService) {
	authEndpoint := &Endpoint{
		attendanceSvc: authSvc,
	}
	r.Route("/api/attendance", func(r chi.Router) {
		r.Use(midleware.Auth.ValidateRoleUser)

		r.Post("/", authEndpoint.checkIn)
		r.Get("/", authEndpoint.getListAttendance)
		r.Put("/", authEndpoint.updateTime)
		r.Put("/check-out", authEndpoint.checkOut)

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

	err := e.attendanceSvc.CheckIn(ctx, int64(userIDCtx), &req)
	if err != nil {
		log.Printf("failed to check in : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}

func (e *Endpoint) checkOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CheckOutRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)

	err := e.attendanceSvc.CheckOut(ctx, int64(userIDCtx), &req)
	if err != nil {
		log.Printf("failed to check out : %s \n", err)
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

	userIDCtx := int64(ctx.Value("UserID").(int32))

	res, err := e.attendanceSvc.GetListAttendanceByUserID(ctx, userIDCtx)
	if err != nil {
		log.Printf("failed to get list attendance : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

func (e *Endpoint) updateTime(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.UpdateTimeAttendanceRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := int64(ctx.Value("UserID").(int32))

	err := e.attendanceSvc.UpdateCheckInTime(ctx, userIDCtx, &req)
	if err != nil {
		log.Printf("failed to update check in time : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
