package internal

func NewErrorWithStatus(status int, err error) error {
	return WithStatusError{
		err:    err,
		status: status,
	}
}

type WithStatusError struct {
	err    error
	status int
}

func (e WithStatusError) Error() string {
	return e.err.Error()
}

func (e WithStatusError) Status() int {
	return e.status
}
