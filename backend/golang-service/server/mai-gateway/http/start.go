package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"

	"manso.live/backend/golang-service/pkg/log"
	"manso.live/backend/golang-service/pkg/util/regutil"
	"manso.live/backend/golang-service/pkg/util/signalutil"
	"manso.live/backend/golang-service/server/mai-gateway/config"
)

func StartHttpServer(cfg *config.Config) {
	server := &Server{
		Config: cfg,
	}

	ch := make(chan struct{})
	go signalutil.SignalHandler(ch, server.Shutdown)

	if err := server.ListenAndServe(); err != nil && !signalutil.IgnoreError(err) {
		log.Error("Failed to start artifacts maven registry server", zap.Error(err))
	}
	<-ch
	log.Info("mai-gateway http server terminated")
}

type Server struct {
	Config *config.Config

	httpServer *http.Server
}

func (s *Server) ListenAndServe() error {
	handler := mux.NewRouter()
	handler.Use(regutil.InitCtxStore)
	handler.Use(regutil.InitLogger)

	healthHandler := handler.PathPrefix("/health").Subrouter()
	healthHandler.
		Methods(http.MethodGet, http.MethodHead).
		Subrouter().
		HandleFunc("/check", s.HealthCheck)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Http.Port),
		Handler: handler,
	}
	log.Info("mai-gateway http server is listening on port", zap.String("port", s.Config.Http.Port))
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	log.Info("Shutting down mai-gateway ...")
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Error("Shutting down mai-gateway failed", zap.Error(err))
	}
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {

}
