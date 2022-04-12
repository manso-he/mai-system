package signalutil

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"syscall"

	"manso.live/backend/golang-service/pkg/util/stringutil"
)

func IgnoreError(err error) bool {
	if err == io.EOF || err == syscall.EPIPE ||
		err == http.ErrServerClosed ||
		stringutil.ContainsAnyString(
			err.Error(),
			"broken pipe",
			"use of closed network connection",
			"connection reset by peer",
			"Invalid protobuf byte sequence",
			"the stream has been done",
		) {
		return true
	}
	if st, ok := status.FromError(err); ok && st.Code() == codes.Canceled {
		return true
	}
	return false
}
