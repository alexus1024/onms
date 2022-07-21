package memory_test

import (
	"sync"
	"testing"
	"time"

	"github.com/alexus1024/onms/internal/models"
	"github.com/alexus1024/onms/internal/storage/memory"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParallelUse(t *testing.T) {
	log := logrus.New().WithField("test", t.Name())
	repo := memory.NewMemoryRepo()

	// 100 goroutines writes 10 entries each
	// another 100 goroutines reads content, 10 tries each
	// no panics or errors and 1000 entries in storage expected in the result

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			m := &models.CapturedData{MachineID: models.MachineID(num + 1)}
			for writeNum := 0; writeNum < 10; writeNum++ {
				err := repo.SaveRecord(m)
				require.NoError(t, err)
				log.WithField("worker", num).Info("inserted")
				time.Sleep(time.Microsecond)
			}
		}(i)

		wg.Add(1)
		go func() {
			defer wg.Done()
			for writeNum := 0; writeNum < 10; writeNum++ {
				d, err := repo.GetAllRecords()
				require.NoError(t, err)
				log.WithField("data_len", len(d)).Info("read")
				time.Sleep(time.Millisecond)
			}
		}()

	}

	wg.Wait()

	allRecords, err := repo.GetAllRecords()
	require.NoError(t, err)
	assert.Len(t, allRecords, 1000)
}
