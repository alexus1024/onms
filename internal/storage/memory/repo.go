package memory

import (
	"fmt"
	"sync"

	"github.com/alexus1024/onms/internal/models"
)

type MemoryRepo struct {
	muMap   sync.RWMutex
	storage map[models.MachineID]*models.CapturedData
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

	r.storage[model.MachineID] = &modelCopy

	return nil
}

func (r *MemoryRepo) GetAllRecords() ([]*models.CapturedData, error) {
	r.muMap.RLock()
	defer r.muMap.RUnlock()

	ret := make([]*models.CapturedData, 0, len(r.storage))

	for _, d := range r.storage {
		ret = append(ret, d)
	}

	return ret, nil
}
