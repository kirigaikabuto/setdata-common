package setdata_common

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	"strings"
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
	if string(responseData.Body) != "null" {
		middleError := &MiddleError{}
		err = json.Unmarshal(responseData.Body, &middleError)
		if err != nil && !strings.Contains(err.Error(), "json: cannot unmarshal array") {
			return err
		}
		if middleError.Code != 0 && middleError.Message != "" {
			return middleError
		}
		err = json.Unmarshal(responseData.Body, &response)
		if err != nil {
			return err
		}
	}
	return nil
}
