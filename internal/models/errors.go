package models

// InputRelatedError describes situations worth of 4xx HTTP return codes
type InputRelatedError struct {
	msg             string
	err             error
	suggestedStatus int
}

func (e InputRelatedError) Error() string {
	return e.msg
}

func (e InputRelatedError) Unwrap() error {
	return e.err
}

func (e InputRelatedError) SuggestedStatus() int {
	return e.suggestedStatus
}

func NewInputRelatedError(msg string, err error) error {
	return InputRelatedError{msg, err, 0}
}

func NewInputRelatedErrorWithStatus(msg string, err error, status int) error {
	if status < 400 || status >= 500 {
		panic("InputRelatedError supports 4xx errors only")
	}
	return InputRelatedError{msg, err, status}
}
