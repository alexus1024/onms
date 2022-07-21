package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexus1024/onms/internal/models"
)

func HandlerCapture(w http.ResponseWriter, r *http.Request, actx *AppContext) error {
	contentType := r.Header.Get(ContentType)
	if contentType != ContentTypeJson {
		return models.NewInputRelatedErrorWithStatus(
			"unsupported content type "+contentType,
			nil,
			http.StatusUnsupportedMediaType,
		)
	}

	jd := json.NewDecoder(r.Body)

	model := &models.CapturedData{}

	err := jd.Decode(model)
	if err != nil {
		return models.NewInputRelatedError("can not decode input", err)
	}

	savedJson, err := json.Marshal(model)
	if err != nil {
		return models.NewInputRelatedError("can not check input by marshalling it", err)
	}

	err = actx.Repo.SaveRecord(model)
	if err != nil {
		return fmt.Errorf("save to storage: %w", err)
	}

	actx.Log.WithField("data", models.JSONRawMessage(savedJson)).Info("new record added")

	w.WriteHeader(http.StatusCreated)

	return nil
}
