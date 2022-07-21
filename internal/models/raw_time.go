package models

import (
	"fmt"
	"strings"
	"time"
)

const CustomApiTimeFormat1 = "Mon 2006-01-02 15:04:05"
const RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"

// allowedApiTimeFormats - here all known allowed time formats are listed.
var allowedApiTimeFormats = []string{
	time.RFC3339,
	CustomApiTimeFormat1,
}

// RawTime represents time in all allowed formats for API:
// * "2022-04-21T19:25:43.219Z"
// * "Wed 2021-07-28 14:16:27"
type RawTime time.Time

func (t *RawTime) UnmarshalJSON(value []byte) error {
	valueStr := string(value)
	valueStr = strings.Trim(valueStr, `"`)

	errorAcc := strings.Builder{}
	for _, format := range allowedApiTimeFormats {
		parsedTime, err := time.Parse(format, valueStr)
		if err == nil {
			*t = RawTime(parsedTime)

			return nil
		}

		errorAcc.WriteString(err.Error())
		errorAcc.WriteString("; ")
	}

	return fmt.Errorf("parse RawTime: %s", errorAcc.String())
}

func (t RawTime) MarshalJSON() ([]byte, error) {
	jsonStr := fmt.Sprintf(`"%s"`, time.Time(t).Format(RFC3339Milli))
	return []byte(jsonStr), nil
}
