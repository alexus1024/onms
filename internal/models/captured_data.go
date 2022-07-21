package models

import "github.com/google/uuid"

// CapturedData represents a report from machine as got from API.
type CapturedData struct {
	MachineID    MachineID         `json:"machineId"`
	Stats        CapturedDataStats `json:"stats"`
	LastLoggedIn string            `json:"lastLoggedIn"`
	SysTime      RawTime           `json:"sysTime"`
}

type CapturedDataStats struct {
	CPUTemp      RawTemperature `json:"cpuTemp,omitempty"`
	FanSpeed     float64        `json:"fanSpeed,omitempty"`
	HDDSpace     float64        `json:"HDDSpace,omitempty"` // nolint:tagliatelle // name by specs
	InternalTemp RawTemperature `json:"internalTemp,omitempty"`
}

// CapturedDataStorage represents storage model for a report.
type CapturedDataStorage struct {
	Id uuid.UUID `json:"id"`
	CapturedData
}
