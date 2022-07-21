package models

type MachineID int

func (id MachineID) IsEmpty() bool {
	return id == 0
}

// TODO: split db model with ID
type CapturedData struct {
	MachineID MachineID `json:"machineId"`
	Stats     struct {
		CPUTemp      RawTemperature `json:"cpuTemp"`
		FanSpeed     float64        `json:"fanSpeed"`
		HDDSpace     float64        `json:"HDDSpace"`
		InternalTemp RawTemperature `json:"internalTemp"`
	} `json:"stats"`
	LastLoggedIn string  `json:"lastLoggedIn"`
	SysTime      RawTime `json:"sysTime"`
}

// TODO: JSON parsers for Raw types
