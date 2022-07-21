package models

import "github.com/google/uuid"

// CapturedData represents a report from machine as got from API
type CapturedData struct {
	MachineID MachineID `json:"machineId"`
	Stats     struct {
		CPUTemp      RawTemperature `json:"cpuTemp,omitempty"`
		FanSpeed     float64        `json:"fanSpeed,omitempty"`
		HDDSpace     float64        `json:"HDDSpace,omitempty"`
		InternalTemp RawTemperature `json:"internalTemp,omitempty"`
	} `json:"stats"`
	LastLoggedIn string  `json:"lastLoggedIn"`
	SysTime      RawTime `json:"sysTime"`
}

// CapturedDataStorage represents storage model for a report
type CapturedDataStorage struct {
	Id uuid.UUID `json:"id"`
	CapturedData
}
