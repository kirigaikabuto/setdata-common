package setdata_common

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

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

func ErrToHttResponse(err error) MiddleError {
	var result MiddleError
	switch err.(type) {
	case MiddleError:
		result = err.(MiddleError)
	default:
		result = NewMiddleError(err, 500, 500)
	}
	return result
}

func ErrToAmqpResponse(err error) *amqp.Message {
	se := ErrToHttResponse(err)
	data, _ := json.Marshal(se)
	return &amqp.Message{Body: data}
}
