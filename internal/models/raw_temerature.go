package models

import (
	"fmt"
	"strconv"
	"strings"
)

// RawTemperature represents temperature's value
// It allows temperature to have optional 'c' in the end.
// NOTE: maybe 'f' is possible too? Could convert to Celsius in this case.
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

func (t RawTemperature) MarshalJSON() ([]byte, error) {
	jsonStr := strconv.FormatFloat(float64(t), 'f', -1, 64)
	return []byte(jsonStr), nil
}
