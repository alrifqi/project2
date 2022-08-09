package utils

// CustomError is type for custom error for this application
// CustomError used for handle error  in app & give response based on CustomError StatusCode
type CustomError struct {
	StatusCode int
	Msg        string
}

// Func for returning error message for CustomError
func (e *CustomError) Error() string {
	return e.Msg
}

// Func for returning error statuscode for CustomError
func (e *CustomError) ErrorStatusCode() int {
	return e.StatusCode
}
