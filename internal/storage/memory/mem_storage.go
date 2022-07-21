package memory

import (
	"fmt"
	"sync"

	"github.com/alexus1024/onms/internal/models"
	"github.com/alexus1024/onms/internal/storage"
	"github.com/google/uuid"
)

var interfaceCheck storage.Repo = &MemoryRepo{}

type MemoryRepo struct {
	muMap   sync.RWMutex
	storage map[uuid.UUID]*models.CapturedData
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		storage: make(map[uuid.UUID]*models.CapturedData),
	}
}

func (r *MemoryRepo) SaveRecord(model *models.CapturedData) error {
	if model == nil {
		return fmt.Errorf("model is nil")
	}
	if model.MachineID.IsEmpty() {
		return fmt.Errorf("machine id is empty")
	}

	modelCopy := *model

	r.muMap.Lock()
	defer r.muMap.Unlock()

	r.storage[uuid.New()] = &modelCopy

	return nil
}

func (r *MemoryRepo) GetAllRecords() ([]*models.CapturedDataStorage, error) {
	r.muMap.RLock()
	defer r.muMap.RUnlock()

	ret := make([]*models.CapturedDataStorage, 0, len(r.storage))

	for id, d := range r.storage {
		ret = append(ret, &models.CapturedDataStorage{Id: id, CapturedData: *d})
	}

	return ret, nil
}
