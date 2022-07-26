package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ApiError is the error model for all API responses.
type ApiError struct {
	Message string
}

// WriteToResponse write status code and error content to the response.
func (e ApiError) WriteToResponse(log *logrus.Entry, w http.ResponseWriter, statusCode int) {
	errOutBytes, err := json.Marshal(e)
	if err != nil {
		log.WithError(err).Error("can not marshal http error (input-related)")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(statusCode)

	_, err = w.Write(errOutBytes)
	if err != nil {
		log.WithError(err).Warn("error on write result")
	}
}
