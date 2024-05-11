package attendance

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0801/internal/model"
	error2 "nam_0801/pkg/error"
	"nam_0801/pkg/util/request"
	"nam_0801/pkg/util/response"
	"net/http"
)

type (
	Endpoint struct {
		authSvc AttendanceService
	}

	AttendanceService interface {
		CheckIn(ctx context.Context, req *model.CheckInRequest) error
		//UpdateTime(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse2, error)
		//GetListAttendanceByUserID(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse2, error)
	}
)

func InitAttendanceHandler(r *chi.Mux, authSvc AttendanceService) {
	authEndpoint := &Endpoint{
		authSvc: authSvc,
	}
	r.Route("/api/attendance", func(r chi.Router) {
		r.Post("/check-in", authEndpoint.checkIn)
		//r.Post("/update-time", authEndpoint.updateTime)
		//r.Get("/", authEndpoint.getListAttendance)
	})
}

func (e *Endpoint) checkIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.LoginRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.authSvc.CheckIn(ctx, &req)
	if err != nil {
		log.Printf("failed to check in : %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

//
//func (e *Endpoint) getListAttendance(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	var req model.LoginRequest
//	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
//		log.Printf("read request body error: %s \n", err)
//		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
//		return
//	}
//
//	res, err := e.authSvc.GetListAttendanceByUserID(ctx, &req)
//	if err != nil {
//		log.Printf("failed to get list attendance : %s \n", err)
//		response.Error(w, err)
//		return
//	}
//
//	response.JSON(w, http.StatusCreated, res)
//}
//
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
