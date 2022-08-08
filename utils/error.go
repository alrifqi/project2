package utils

type CustomError struct {
	StatusCode int
	Msg        string
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) ErrorStatusCode() int {
	return e.StatusCode
}
