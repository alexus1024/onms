package models

// InputRelatedError describes situations worth of 4xx HTTP return codes
type InputRelatedError struct {
	msg string
	err error
}

func (e InputRelatedError) Error() string {
	return e.msg
}

func (e InputRelatedError) Unwrap() error {
	return e.err
}

func NewInputRelatedError(msg string, err error) error {
	return InputRelatedError{msg, err}
}
