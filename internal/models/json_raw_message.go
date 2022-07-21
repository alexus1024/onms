package models

import (
	"encoding/json"
	"errors"
)

var ErrNilJsonPointer = errors.New("JSONRawMessage: UnmarshalJSON on nil pointer")

// JSONRawMessage is like a json.RawMessage but Sprint outputs a JSON string for it, not bytes.
// E.g. logrus.WithField("key", JSONRawMessage{}) will print a text with both JSON and Console formatters.
type JSONRawMessage json.RawMessage

func (m JSONRawMessage) String() string {
	return string(m)
}

// MarshalJSON returns m as the JSON encoding of m.
func (m JSONRawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}

	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *JSONRawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return ErrNilJsonPointer
	}

	*m = append((*m)[0:0], data...)

	return nil
}
