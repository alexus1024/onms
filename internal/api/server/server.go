package server

import (
	"net/http"

	"github.com/alexus1024/onms/internal/models"
	"github.com/alexus1024/onms/internal/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	ContentType     = "content-type"
	ContentTypeJson = "application/json"
)

// AppHandler is a http.Handler enriched with application context.
// In case of error it should return go error instead of using Writer.
type AppHandler func(http.ResponseWriter, *http.Request, *AppContext) error

// AppContext stores runtime dependencies for server's handlers
// Not protected from changes, so just do not change it's content in handlers please.
// But you can copy it instead.
type AppContext struct {
	Log  *logrus.Entry
	Repo storage.Repo
}

// GetMux returns a handler configured to process all required operations.
// Sub-handlers require server context to be configured properly.
func GetMux(actx *AppContext) http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/", toHandler(HandlerCapture, actx)).Methods(http.MethodPost).Name("capture")
	r.HandleFunc("/", toHandler(HandlerGetAll, actx)).Methods(http.MethodGet).Name("read all")

	return r
}

// toHandler converts AppHandler to http.HandlerFunc.
// It also processes Golang errors into HTTP errors and thereby
// setups application-wide standard for HTTP errors
func toHandler(ah AppHandler, actx *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		log := actx.Log.WithField("handler", route.GetName())

		// add current route into all handler's logs
		ctxCopy := *actx
		ctxCopy.Log = log

		// call actual handler
		err := ah(w, r, &ctxCopy)

		// convert go errors into API errors
		if err != nil {
			route := mux.CurrentRoute(r)
			log := actx.Log.WithField("route", route.GetName())
			switch te := err.(type) {
			// TODO: reduce code duplication
			case models.InputRelatedError:
				log.WithError(err).Info("api error (input-related)")
				errOut := ApiError{Message: te.Error()}
				status := te.SuggestedStatus()
				if status == 0 {
					status = 400
				}
				errOut.WriteToResponse(log, w, status)
				return

			default:
				log.WithError(err).Error("api error")
				errOut := ApiError{Message: "internal error"}
				errOut.WriteToResponse(log, w, 500)
				return
			}
		}

	}

}
