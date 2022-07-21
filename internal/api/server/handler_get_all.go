package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerGetAll(w http.ResponseWriter, r *http.Request, actx *AppContext) error {
	dbData, err := actx.Repo.GetAllRecords()
	if err != nil {
		return fmt.Errorf("read from storage: %w", err)
	}

	w.Header().Add(ContentType, ContentTypeJson)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(dbData)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return nil
}
