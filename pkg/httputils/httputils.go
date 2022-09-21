package httputils

import (
	"net/http"
)

func NotImplementedHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("Not implemented yet."))
}

// StatusRecorder implements the http.ResponseWriter interface.
// It stores the last status code passed to the underlying response writer
type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (w *StatusRecorder) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
