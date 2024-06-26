package authentication

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0511/internal/model"
	error2 "nam_0511/pkg/error"
	"nam_0511/pkg/util/request"
	"nam_0511/pkg/util/response"
	"net/http"
)

type (
	Endpoint struct {
		authSvc AuthenticationService
	}

	AuthenticationService interface {
		Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	}
)

func InitAuthenticationHandler(r *chi.Mux, authSvc AuthenticationService) {
	authEndpoint := &Endpoint{
		authSvc: authSvc,
	}
	r.Route("/", func(r chi.Router) {
		r.Post("/api/login", authEndpoint.login)
	})
}

// @Summary Login
// @Description Login by email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.LoginRequest true "Login"
// @Success 200 {object} model.LoginResponse
// @Router /api/login [post]
func (e *Endpoint) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.LoginRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.authSvc.Login(ctx, &req)
	if err != nil {
		log.Printf("failed to create user: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
