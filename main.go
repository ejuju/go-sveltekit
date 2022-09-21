package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ejuju/go-sveltekit/pkg/fsutil"
	"github.com/ejuju/go-sveltekit/pkg/httputils"
	"github.com/ejuju/go-sveltekit/pkg/logutil"
	"github.com/ejuju/go-sveltekit/website"
	"github.com/gorilla/mux"
)

type HTTPRouter struct {
	WebsiteHTTPHandler http.Handler
	BackendHTTPHandler http.Handler
}

// This function routes requests to the appropriate handler
// depending if they are for the backend API or the file server (= website files).
func (httpRouter *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api/") {
		httpRouter.WebsiteHTTPHandler.ServeHTTP(w, r)
		return
	}

	httpRouter.BackendHTTPHandler.ServeHTTP(w, r)
}

func main() {
	logger := &logutil.DefaultLogger{}

	// Get HTTP port from environment variable
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Errorf("invalid port number: %w", err))
	}

	// Get website build sub file-system
	websiteFS, err := fs.Sub(website.BuildFS, "build")
	if err != nil {
		panic(fmt.Errorf("unable to get website build sub file-system: %w", err))
	}
	// rawFallbackPage, err := websiteFS.Open("404.html")
	// if err != nil {
	// 	panic(fmt.Errorf("unable to open index.html page: %w", err))
	// }
	// fallbackPage, err := ioutil.ReadAll(rawFallbackPage)
	// if err != nil {
	// 	panic(fmt.Errorf("unable to read index.html page: %w", err))
	// }

	// Print website files for debugging
	err = fsutil.LogFiles(logger, logutil.LogLevelDebug, websiteFS)
	if err != nil {
		panic(fmt.Errorf("unable to log website static files: %w", err))
	}

	// Init website handler
	httpWebsiteHandler := mux.NewRouter()
	httpWebsiteHandler.PathPrefix("/").Handler(http.FileServer(http.FS(websiteFS)))

	// Init backend HTTP router
	httpBackendHandler := mux.NewRouter()
	httpBackendHandler.HandleFunc("/api/v1/", httputils.NotImplementedHandlerFunc)

	// Init HTTP router wrapper
	var httpHandler http.Handler = &HTTPRouter{
		WebsiteHTTPHandler: httpWebsiteHandler,
		BackendHTTPHandler: httpBackendHandler,
	}
	httpHandler = logutil.NewHTTPLoggerMiddleware(logger)(httpHandler)

	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: httpHandler,
		// todo: set defaults
	}

	logger.Log(logutil.LogLevelInfo, "Starting HTTP server on port "+strconv.Itoa(port))
	err = httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
