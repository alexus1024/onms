package storage

import "github.com/alexus1024/onms/internal/models"

type Repo interface {
	SaveRecord(*models.CapturedData) error
	GetAllRecords() ([]*models.CapturedData, error)
}
