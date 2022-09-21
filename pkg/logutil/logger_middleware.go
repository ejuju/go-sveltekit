package logutil

import (
	"net/http"
	"strconv"

	"github.com/ejuju/go-sveltekit/pkg/httputils"
)

func NewHTTPLoggerMiddleware(logger Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &httputils.StatusRecorder{ResponseWriter: w}
			h.ServeHTTP(recorder, r)
			logger.Log(LogLevelInfo, "Served HTTP request: "+strconv.Itoa(recorder.Status)+" "+r.URL.Path)
		})
	}
}
