package server

import (
	"encoding/json"
	"net/http"

	"github.com/alexus1024/onms/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Message string
}

func GetMux(log *logrus.Entry) http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/", handleErrors(HandlerCapture, log)).Methods(http.MethodPost).Name("capture")
	r.HandleFunc("/", handleErrors(HandlerGetAll, log)).Methods(http.MethodGet).Name("read all")

	return r
}

// handleErrors converts Golang erros into HTTP errors and thereby
// setups application-wide standard for HTTP errors
func handleErrors(ah AppHandler, log *logrus.Entry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ah(w, r)
		if err != nil {
			route := mux.CurrentRoute(r)
			log := log.WithField("route", route.GetName())
			switch te := err.(type) {
			// TODO: reduce code dublication
			case models.InputRelatedError:
				log.WithError(err).Info("api error (input-related)")
				errOut := ApiError{Message: te.Error()}
				errOutBytes, err := json.Marshal(errOut)
				if err != nil {
					log.WithError(err).Error("can not marshal http error (input-related)")
					w.WriteHeader(500)
					return
				}
				w.Write(errOutBytes)
				w.WriteHeader(400)
				return

			default:
				log.WithError(err).Error("api error")
				errOut := ApiError{Message: "internal error"}
				errOutBytes, err := json.Marshal(errOut)
				if err != nil {
					log.WithError(err).Error("can not marshal http error")
					w.WriteHeader(500)
					return
				}
				w.Write(errOutBytes)
				w.WriteHeader(500)
				return
			}
		}

	}

}
