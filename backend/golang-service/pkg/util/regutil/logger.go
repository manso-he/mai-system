package regutil

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"manso.live/backend/golang-service/pkg/log"
	"manso.live/backend/golang-service/pkg/util/httputil"
	"manso.live/backend/golang-service/pkg/util/uuidutil"
	"net/http"
	"strings"
	"time"
)

const (
	RequestId = "requestId"
)

type logResponseWriter struct {
	http.ResponseWriter
	Status int

	httpMethod string
}

func (w *logResponseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *logResponseWriter) Write(b []byte) (int, error) {
	if strings.EqualFold(w.httpMethod, http.MethodHead) {
		return len(b), nil
	}
	return w.ResponseWriter.Write(b)
}

func InitLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuidutil.Generate()

		logger := log.With(zap.String(RequestId, requestId))
		r = r.WithContext(context.WithValue(r.Context(), requestIdKey, requestId))
		r = r.WithContext(context.WithValue(r.Context(), loggerKey, logger))

		lrw := &logResponseWriter{ResponseWriter: w, httpMethod: r.Method}

		startTime := time.Now()
		next.ServeHTTP(lrw, r)
		finishTime := time.Now()

		logger.With(
			zap.String("method", r.Method),
			zap.String("host", r.Host),
			zap.String("uri", r.RequestURI),
			zap.String("client_ip", httputil.ClientIP(r)),
			zap.String("user_agent", r.UserAgent()),
			zap.Int("status", lrw.Status),
			zap.String("elapsed_time", fmt.Sprintf("%.4fs", finishTime.Sub(startTime).Seconds())),
		).Info("access log")
	})
}

func Logger(r *http.Request) *zap.Logger {
	return r.Context().Value(loggerKey).(*zap.Logger)
}
