package setdata_common

type MiddleError struct {
	Code         int
	Err          error
	AddInfo      interface{} `json:"add_info,omitempty"`
	UniqueNumber int
}

func (se MiddleError) Error() string {
	return se.Err.Error()
}

func (se MiddleError) Status() int {
	return se.Code
}

func NewMiddleError(err error, statusCode, unique int) MiddleError {
	switch err.(type) {
	case MiddleError:
		return err.(MiddleError)
	default:
		return MiddleError{
			Code:         statusCode,
			Err:          err,
			UniqueNumber: unique,
		}
	}
}