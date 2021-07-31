package setdata_common

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
	"reflect"
)

func AmqpCall(clt amqp.Client, endpoint string, request interface{}, response interface{}) error {
	requestData, err := json.Marshal(request)
	if err != nil {
		return err
	}
	responseData, err := clt.Call(endpoint, amqp.Message{Body: requestData})
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseData.Body, &response)
	if err != nil {
		return err
	}
	fmt.Println(response)
	fmt.Println(reflect.TypeOf(response))
	switch response.(type) {
	case MiddleError:
		return response.(MiddleError)
	default:
		return nil
	}
}
