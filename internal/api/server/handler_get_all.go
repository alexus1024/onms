package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexus1024/onms/internal/storage"
)

func HandlerGetAll(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	repo := storage.GetStorage(ctx)

	dbData, err := repo.GetAllRecords()
	if err != nil {
		return fmt.Errorf("read from storage: %w", err)
	}

	jencoder := json.NewEncoder(w)
	err = jencoder.Encode(dbData)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return nil
}
