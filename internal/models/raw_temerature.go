package models

import (
	"fmt"
	"strconv"
	"strings"
)

// RawTemperature represents temperature's value
// It allows temperature to have optional 'c' in the end
type RawTemperature float64

func (t *RawTemperature) UnmarshalJSON(value []byte) error {
	valueStr := string(value)
	valueStr = strings.Trim(valueStr, `"c`)

	parsedValue, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("parse RawTemperature: %w", err)
	}

	*t = RawTemperature(parsedValue)

	return nil
}
