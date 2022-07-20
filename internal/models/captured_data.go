package models

import "time"

type RawTemperature float64
type RawTime time.Time
type MachineID int

type CapturedData struct {
	MachineID MachineID `json:"machineId"`
	Stats     struct {
		CPUTemp      RawTemperature `json:"cpuTemp"`
		FanSpeed     float64        `json:"fanSpeed"`
		HDDSpace     float64        `json:"HDDSpace"`
		InternalTemp float64        `json:"internalTemp"`
	} `json:"stats"`
	LastLoggedIn string  `json:"lastLoggedIn"`
	SysTime      RawTime `json:"sysTime"`
}

// TODO: JSON parsers for Raw types
