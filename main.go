package main

import (
	"context"
	"errors"
	"log"
	authentication2 "nam_0801/internal/endpoint/authentication"
	user3 "nam_0801/internal/endpoint/user"
	"nam_0801/internal/repo/repo/user"
	"nam_0801/internal/service/authentication"
	user2 "nam_0801/internal/service/user"
	"nam_0801/pkg/config"
	"nam_0801/pkg/midleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"nam_0801/pkg/db/postgres"
)

func main() {
	startedAt := time.Now()
	defer func() {
		log.Printf("application stopped after %s\n", time.Since(startedAt))
	}()

	conf, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}

	globalCtx, glbCtxCancel := context.WithCancel(context.Background())

	httpSrv, err := initHTTPServer(globalCtx, conf)
	if err != nil {
		log.Panicf("failed to init http server %s \n", err)
	}

	go func() {
		log.Printf("starting HTTP server at: %s\n", conf.Server.Address)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("failed to start HTTP server: %s \n", err)
		}
	}()

	// Keep the application running until signals trapped
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("%s signal trapped. Stopping application", <-sigChan)

	glbCtxCancel()
	// First terminate the HTTP gateway
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout)
	defer shutdownCtxCancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to gracefully shutdown the HTTP gateway server: %s\n", err)
	} else {
		log.Println("HTTP gateway server stopped gracefully")
	}
}

func initHTTPServer(ctx context.Context, conf *configs.Config) (httpServer *http.Server, err error) {
	r := chi.NewRouter()

	// create endpoint here
	r.Get("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	dbConn, err := postgres.ConnectDatabase(conf.Postgres)
	if err != nil {
		log.Panicf("failed to connect database:: %s \n", err)
		return
	}

	midleware.Auth.Cfg = conf.Server.Token

	// repository
	userRepo := user.NewPostgresRepository(dbConn)

	// service
	userService := user2.NewUserService(dbConn, userRepo)
	authService := authentication.NewAuthenticationService(dbConn, conf.Server.Token, userRepo)

	// handler
	user3.InitUserHandler(r, userService)
	authentication2.InitAuthenticationHandler(r, authService)

	return &http.Server{
		Addr:         conf.Server.Address,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
		Handler:      r,
	}, nil
}
