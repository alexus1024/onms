package models

type MachineID int

func (id MachineID) IsEmpty() bool {
	return id == 0
}
