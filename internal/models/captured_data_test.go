package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alexus1024/onms/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

const (
	json1 = `{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90,
        "fanSpeed": 400,
        "HDDSpace": 800
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}`
	json2 = `{
    "machineId": 4444,
    "stats": {
        "cpuTemp": 78,
        "fanSpeed": 500,
        "HDDSpace": 100,
        "internalTemp": 23
    },
    "lastLoggedIn": "admin/Ian",
    "sysTime": "2022-04-21T19:25:43.219Z"
}`
	json3 = `{
    "machineId": 61616,
    "stats": {
        "cpuTemp": "78c",
        "fanSpeed": 500,
        "HDDSpace": 100,
        "internalTemp": 23
    },
    "lastLoggedIn": "admin/Tim",
    "sysTime": "Wed 2021-07-28 14:16:27"
}`
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc         string
		json         string
		expMachineId models.MachineID
		expCpuTemp   float64
		expSysTime   time.Time
	}{
		{
			desc:         "case 1",
			json:         json1,
			expMachineId: 12345,
			expCpuTemp:   90,
			expSysTime:   time.Date(2022, 4, 23, 18, 25, 43, 511000000, time.UTC),
		},
		{
			desc:         "case 2",
			json:         json2,
			expMachineId: 4444,
			expCpuTemp:   78,
			expSysTime:   time.Date(2022, 4, 21, 19, 25, 43, 219000000, time.UTC),
		},
		{
			desc:         "case 3",
			json:         json3,
			expMachineId: 61616,
			expCpuTemp:   78,
			expSysTime:   time.Date(2021, 7, 28, 14, 16, 27, 0, time.UTC),
		},
	}
	for _, tC := range testCases {
		tC := tC // closure

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := &models.CapturedData{}
			err := json.Unmarshal([]byte(tC.json), model)
			require.NoError(t, err)
			assert.Equal(t, tC.expMachineId, model.MachineID)
			assert.Equal(t, tC.expCpuTemp, float64(model.Stats.CPUTemp))
			assert.Equal(t, tC.expSysTime, time.Time(model.SysTime))
		})
	}
}

func TestMarshal(t *testing.T) {
	t.Parallel()

	demoTZ, err := time.LoadLocation("Asia/Yekaterinburg") // GMT+5

	require.NoError(t, err)

	model := &models.CapturedData{
		MachineID: 1,
		Stats: models.CapturedDataStats{
			CPUTemp:  2,
			FanSpeed: 3,
		},
		LastLoggedIn: "4",
		SysTime:      models.RawTime(time.Date(2022, 1, 2, 3, 4, 5, 6, demoTZ)),
	}

	jsonResult, err := json.Marshal(model)
	require.NoError(t, err)

	gjson.ValidBytes(jsonResult)
	assert.Equal(t, float64(1), gjson.GetBytes(jsonResult, "machineId").Num)
	assert.Equal(t, float64(2), gjson.GetBytes(jsonResult, "stats.cpuTemp").Num)
	assert.Equal(t, float64(3), gjson.GetBytes(jsonResult, "stats.fanSpeed").Num)
	assert.Equal(t, "4", gjson.GetBytes(jsonResult, "lastLoggedIn").Str)
	assert.Equal(t, "2022-01-02T03:04:05+05:00", gjson.GetBytes(jsonResult, "sysTime").Str) // 6 nanoseconds was cut
}
