package setdata_common

import (
	"encoding/json"
	"errors"
	"github.com/djumanoff/amqp"
)

var (
	ErrCurrentAclResourceNotExist = NewMiddleError(errors.New("current acl resource not exist"), 404, 100)
	ErrCurrentAclActionNotExist = NewMiddleError(errors.New("current acl action not exist"), 404, 101)
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

func ErrToHttpResponse(err error) MiddleError {
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
	se := ErrToHttpResponse(err)
	data, _ := json.Marshal(se)
	return &amqp.Message{Body: data}
}
