package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type RouterError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err RouterError) Error() string {
	return fmt.Sprintf("error %d: %s", err.Code, err.Message)
}

func MarshalResponse(status int, headers map[string]string, data interface{}) (
	events.APIGatewayProxyResponse,
	error,
) {
	b, err := json.Marshal(data)
	if err != nil {
		status = http.StatusInternalServerError
		b = []byte(`{"code":500,"message":"the server has encountered an unexpected error"}`)
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json; charset=UTF-8"

	return events.APIGatewayProxyResponse{
		StatusCode:      status,
		IsBase64Encoded: false,
		Headers:         headers,
		Body:            string(b),
	}, nil
}

var ExposeServerErrors = true

func HandleError(err error) (events.APIGatewayProxyResponse, error) {
	var rErr RouterError
	if !errors.As(err, &rErr) {
		rErr = RouterError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if rErr.Code >= 500 && !ExposeServerErrors {
		rErr.Message = http.StatusText(rErr.Code)
	}

	return MarshalResponse(rErr.Code, nil, rErr)
}
