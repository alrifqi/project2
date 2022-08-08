package utils

type CustomerError struct {
	StatusCode int
	Msg        string
}

func (e *CustomerError) Error() string {
	return e.Msg
}

func (e *CustomerError) ErrorStatusCode() int {
	return e.StatusCode
}
