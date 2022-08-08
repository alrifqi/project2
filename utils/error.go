package utils

type DataNotFoundError struct {
	StatusCode int
	Msg        string
}

type InvalidDataError struct {
	StatusCode int
	Msg        string
}

type CustomerError struct {
	StatusCode int
	Msg        string
}

func (e *DataNotFoundError) Error() string {
	return e.Msg
}

func (e *DataNotFoundError) ErrorStatusCode() int {
	return e.StatusCode
}

func (e *CustomerError) Error() string {
	return e.Msg
}

func (e *CustomerError) ErrorStatusCode() int {
	return e.StatusCode
}
