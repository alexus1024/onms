package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexus1024/onms/internal/models"
	"github.com/alexus1024/onms/internal/storage"
)

func HandlerCapture(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	repo := storage.GetStorage(ctx)

	jd := json.NewDecoder(r.Body)

	model := &models.CapturedData{}
	err := jd.Decode(model)
	if err != nil {
		return models.NewInputRelatedError("can not decode input", err)
	}

	err = repo.SaveRecord(model)
	if err != nil {
		return fmt.Errorf("save to storage: %w", err)
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
